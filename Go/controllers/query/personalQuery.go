package query

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	ebook "yisinnft.org/m/v2/contracts"
)

type QueryPersonalController struct {
	Instance *ebook.YiSinEBook
	DB       *mongo.Client
}

func (con QueryPersonalController) GetRentBookEndTime(ctx *gin.Context) {
	idstr := ctx.Param("id")
	idBigInt, ok := new(big.Int).SetString(idstr, 10)
	//	Test id correct
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	addressStr := ctx.Param("address")
	address := common.HexToAddress(addressStr)
	index, err := con.Instance.RenterRentInfoIndex(nil, address, idBigInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Index_error": err})
	}

	info, err := con.Instance.RentInfos(nil, idBigInt, index)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"rent_info_index": err})
	}

	endTime := info.EndTime
	now := time.Now().Unix()
	duration := endTime.Sub(endTime, big.NewInt(now))
	durationToSecond := new(big.Int).Set(duration)
	secondInt64 := durationToSecond.Int64()

	days := secondInt64 / (24 * 60 * 60)
	hours := (secondInt64 % (24 * 60 * 60)) / (60 * 60)
	minutes := (secondInt64 % (60 * 60)) / 60
	seconds := secondInt64 % 60

	ctx.JSON(http.StatusOK, gin.H{
		"remaining_in_second": duration,
		"seconds_remaining":   seconds,
		"minutes_remaining":   minutes,
		"hours_remaining":     hours,
		"days_remaining":      days,
	})
}

func (con QueryPersonalController) GetPersonalRentedBook(ctx *gin.Context) {
	addressStr := ctx.Param("address")

	//	Use moralis api to get personal nft
	url := fmt.Sprintf("https://deep-index.moralis.io/api/v2.2/%s/nft?chain=sepolia&format=decimal&token_addresses%%5B0%%5D=0x790e48C4F57F4415b9Aed58157A6A8436ea094A6&media_items=false", addressStr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjAxMzA5NzA1LTA0YjYtNDRkZS05YThlLTBjMzU5ZjA0ZGZhOCIsIm9yZ0lkIjoiMzg3NjI5IiwidXNlcklkIjoiMzk4Mjk2IiwidHlwZUlkIjoiNTcxMTgwZjYtYWUxZS00ZGU4LTk3NGQtYTJmMDljNDUwM2VlIiwidHlwZSI6IlBST0pFQ1QiLCJpYXQiOjE3MTI5MTI4MTYsImV4cCI6NDg2ODY3MjgxNn0.u6VKPRrHScpsHzf83hht5-E1UKKDz3DlQ6BMlayaFLg")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// Use json.Unmarshal, analyze body(json) to result(gin.H)
	var result gin.H
	json.Unmarshal(body, &result)

	// extract token_id
	items := result["result"].([]interface{})
	var extractedItems []int64
	for _, item := range items {
		itemMap := item.(map[string]interface{})
		tokenID := itemMap["token_id"].(string)
		tokenid, _ := strconv.ParseInt(tokenID, 10, 64)
		extractedItems = append(extractedItems, tokenid)
	}

	db := con.DB.Database("ebook")

	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": "Failing when search collections",
		})
		return
	}

	bookChannel := make(chan Book)
	books := make([]Book, 0)
	go func() {
		for book := range bookChannel {
			books = append(books, book)
		}
	}()

	var wg sync.WaitGroup

	for _, collName := range collections {
		wg.Add(1)

		go func(collName string) {
			defer wg.Done()

			coll := db.Collection(collName)
			//	Find tokenId in extractedItems
			filter := bson.M{"tokenId": bson.M{"$in": extractedItems}}

			cur, _ := coll.Find(context.TODO(), filter)
			if cur == nil {
				return
			}

			for cur.Next(context.TODO()) {
				var result Book
				err = cur.Decode(&result)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error": "Fail occur while decoding result",
					})
					return
				}
				bookChannel <- result
			}

			cur.Close(context.TODO())
		}(collName)
	}
	wg.Wait()
	close(bookChannel)

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})

}

func (con QueryPersonalController) GetPersonalPublish(ctx *gin.Context) {
	address := ctx.Param("address")
	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": "Failing when search collections",
		})
		return
	}

	var wg sync.WaitGroup
	bookChannel := make(chan Book)
	books := make([]Book, 0)

	//	Start a goroutine as consumer
	go func() {
		for book := range bookChannel {
			books = append(books, book)
		}
	}()
	//	A goroutine corresponds to a collection
	for _, collName := range collections {
		//	Add one wait group
		wg.Add(1)
		go func(collName string) {
			//	Decrease one when function be solved
			defer wg.Done()
			coll := db.Collection(collName)
			filter := bson.M{"uploader": address}

			cur, _ := coll.Find(context.TODO(), filter)
			if cur == nil {
				return
			}

			for cur.Next(context.TODO()) {
				var book Book
				if err := cur.Decode(&book); err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error": "Error occur while decoding data",
					})
					return
				}
				bookChannel <- book
			}
		}(collName)
	}
	//	Wait until counter is zero
	wg.Wait()
	close(bookChannel)

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (con QueryPersonalController) GetBookFile(ctx *gin.Context){

}

func (con QueryPersonalController) VerifySignatureMiddleWare(ctx *gin.Context) {
	signature := ctx.Param("signature")
	publicKey := ctx.Param("address")

	publicKeyByte, err := hexutil.Decode(publicKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error occur while decode public key",
		})
		return
	}

	signatureByte, err := hexutil.Decode(signature)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error occur while decode signature",
		})
		return
	}

	data := []byte(`Welcome to YiSin ebook store!
	
Click to verify that you own this wallet and have control over it.

YiSin ebook (https://yisinnft.org/ebook) need to confirm whether you have the permission to read the e-book file.

This request will not trigger a blockchain transaction or cost any gas fees.`)

	hash := crypto.Keccak256Hash(data)

	signatureNoRecoverID := signatureByte[:len(signatureByte)-1]

	verified := crypto.VerifySignature(publicKeyByte, hash.Bytes(), signatureNoRecoverID)

	if !verified {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "Signature verify fail",
		})
		return
	}
	ctx.Next()
}

func (con QueryPersonalController) CheckPermissionToAccessFileMiddleWare(ctx *gin.Context) {
	publicKey := ctx.Param("address")
	tokenIdStr := ctx.Param("id")

	address := common.HexToAddress(publicKey)

	tokenIdBigInt := new(big.Int)
	tokenIdBigInt, ok := tokenIdBigInt.SetString(tokenIdStr, 10)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Transform tokenId error",
		})
		return
	}

	addressHaveTokenId, err := con.Instance.IsAddressHaveTokenId(nil, address, tokenIdBigInt)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error":   "Error occur while check blockchain information",
			"message": err.Error(),
		})
		return
	}

	if addressHaveTokenId{
		ctx.Next()
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "You not have this book's NFT, please borrow it first",
		})
		return
	}
}
