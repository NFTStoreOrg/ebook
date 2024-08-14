package query

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	ebook "yisinnft.org/m/v2/contracts"
	"yisinnft.org/m/v2/models"
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
		ctx.String(http.StatusBadRequest, "Invalid ID")
	}
	addressStr := ctx.Param("address")
	address := common.HexToAddress(addressStr)
	index, err := con.Instance.RenterRentInfoIndex(nil, address, idBigInt)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	info, err := con.Instance.RentInfos(nil, idBigInt, index)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
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
	url := fmt.Sprintf("https://deep-index.moralis.io/api/v2.2/%s/nft?chain=sepolia&format=decimal&token_addresses%%5B0%%5D=0xD658Ca5061B4e5bAbFAA49c70A52033dC1f98a78&media_items=false", addressStr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjAxMzA5NzA1LTA0YjYtNDRkZS05YThlLTBjMzU5ZjA0ZGZhOCIsIm9yZ0lkIjoiMzg3NjI5IiwidXNlcklkIjoiMzk4Mjk2IiwidHlwZUlkIjoiNTcxMTgwZjYtYWUxZS00ZGU4LTk3NGQtYTJmMDljNDUwM2VlIiwidHlwZSI6IlBST0pFQ1QiLCJpYXQiOjE3MTI5MTI4MTYsImV4cCI6NDg2ODY3MjgxNn0.u6VKPRrHScpsHzf83hht5-E1UKKDz3DlQ6BMlayaFLg")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
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

	collections, err := db.ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
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

			cur, _ := coll.Find(context.Background(), filter)
			if cur == nil {
				return
			}

			for cur.Next(context.TODO()) {
				var result Book
				err = cur.Decode(&result)
				if err != nil {
					ctx.String(http.StatusInternalServerError, err.Error())
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

	//	Marshal data to json and store it in redis.
	jsonData, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//	Set local random source
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	//	Set min and max minutes
	min := 5
	max := 10
	//	Calculate expiration time
	expirationTime := (rng.Intn(max-min+1) + min) * int(time.Minute)

	//	Set data in redis
	_, err = models.RedisClient.Set(context.Background(), "personal_rented_"+addressStr, string(jsonData), time.Duration(expirationTime)).Result()
	if err != nil {
		log.Fatal("Error while set newest book in redis: ", err.Error())
	}

}

func (con QueryPersonalController) GetPersonalPublish(ctx *gin.Context) {
	address := ctx.Param("address")
	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.Background(), bson.M{})

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
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

			cur, _ := coll.Find(context.Background(), filter)
			if cur == nil {
				return
			}

			for cur.Next(context.Background()) {
				var book Book
				if err := cur.Decode(&book); err != nil {
					ctx.String(http.StatusInternalServerError, err.Error())
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

	//	Marshal data to json and store it in redis.
	jsonData, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//	Set local random source
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	//	Set min and max minutes
	min := 5
	max := 10
	//	Calculate expiration time
	expirationTime := (rng.Intn(max-min+1) + min) * int(time.Minute)

	//	Set data in redis
	_, err = models.RedisClient.Set(context.Background(), "personal_publish_"+address, string(jsonData), time.Duration(expirationTime)).Result()
	if err != nil {
		log.Fatal("Error while set newest book in redis: ", err.Error())
	}
}

func (con QueryPersonalController) GetBookFile(ctx *gin.Context) {
	idStr := ctx.Param("id")

	// Attempt to find file in given path
	match, err := filepath.Glob(filepath.Join("./static/bookfile", idStr+".*"))
	if err != nil {
		if os.IsNotExist(err) {
			ctx.String(http.StatusBadRequest, "Invalid path or file not found")
		} else {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	}
	if len(match) == 0 {
		ctx.String(http.StatusNotFound, "File not found")
	}

	//	Return file
	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition")
	ctx.Header("Content-Disposition", "attachment; filename="+filepath.Base(match[0]))	//	filepath.Base: the last router, before . string
	ctx.File(match[0])
}

func (con QueryPersonalController) AddressHaveRentedBook(ctx *gin.Context) {
	addressStr := ctx.Param("address")
	idStr := ctx.Param("id")

	address := common.HexToAddress(addressStr)

	tokenId := new(big.Int)
	tokenId, ok := tokenId.SetString(idStr, 10)
	if !ok {
		ctx.String(http.StatusBadRequest, "Transform tokenId fail")
	}

	isRented, err := con.Instance.IsAddressHaveTokenId(nil, address, tokenId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
	}

	if !isRented {
		ctx.JSON(http.StatusOK, gin.H{
			"rentStatus": false,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"rentStatus": true,
		})
	}

}

func (con QueryPersonalController) VerifySignatureMiddleWare(ctx *gin.Context) {
	signature := ctx.Param("signature")
	publicKey := ctx.Param("address")

	signatureByte, err := hexutil.Decode(signature)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	if len(signatureByte) != 65 {
		ctx.String(http.StatusBadRequest, "Signature length must have 65 bytes")
	}

	v, r, s := signatureByte[64], signatureByte[:32], signatureByte[32:64]
	if v != 0 && v != 1 {
		v -= 27
	}
	data := []byte(`Welcome to YiSin ebook store!

Click to verify that you own this wallet and have control over it.

YiSin ebook (https://yisinnft.org/ebook) need to confirm whether you have the permission to read the e-book file.

This request will not trigger a blockchain transaction or cost any gas fees.`)

	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data))
	hash := crypto.Keccak256Hash([]byte(prefix), []byte(data))

	//	Recovery public key from signature.
	recoveredPubKey, err := crypto.SigToPub(hash.Bytes(), append(r, append(s, v)...))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	//	Change public key to address
	recoveredAddress := strings.ToLower(crypto.PubkeyToAddress(*recoveredPubKey).Hex())

	if recoveredAddress != publicKey {
		ctx.String(http.StatusForbidden, "Signature verify fail")
	}
}

func (con QueryPersonalController) CheckPermissionToAccessFileMiddleWare(ctx *gin.Context) {
	publicKey := ctx.Param("address")
	tokenIdStr := ctx.Param("id")

	address := common.HexToAddress(publicKey)

	tokenIdBigInt := new(big.Int)
	tokenIdBigInt, ok := tokenIdBigInt.SetString(tokenIdStr, 10)
	if !ok {
		ctx.String(http.StatusBadRequest, "Transform tokenId error")
	}

	addressHaveTokenId, err := con.Instance.IsAddressHaveTokenId(nil, address, tokenIdBigInt)
	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
	}

	if addressHaveTokenId {
		ctx.Next()
	} else {
		ctx.String(http.StatusForbidden, "You not have this book's NFT, please borrow it first")
	}
}
