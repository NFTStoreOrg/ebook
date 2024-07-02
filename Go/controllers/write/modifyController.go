package write

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yisinnft.org/m/v2/controllers/query"
)

type ModifyController struct {
	DB   *mongo.Client
	Book *query.Book
}

func (con ModifyController) AdjustBookInformation(ctx *gin.Context) {
	id := ctx.Param("id")

	//	Use json format receive data
	var updateData map[string]interface{}
	if err := ctx.BindJSON(&updateData); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	for _, collectionName := range collections {
		collection := db.Collection(collectionName)
		filter := bson.M{"tokenId": id}
		update := bson.M{"$set": updateData}

		result, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		//	If update one data, break for loop
		if result.ModifiedCount > 0 {
			ctx.JSON(http.StatusOK, gin.H{"message": "Update successfully"})
			return
		}
	}

	ctx.String(http.StatusNotFound, "TokenId not found")
}

func (con ModifyController) VerifySignatureMiddleWare(ctx *gin.Context) {
	idstr := ctx.Param("id")
	signature := ctx.Param("signature")

	signatureByte, err := hexutil.Decode(signature)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	data := []byte(`Welcome to YiSin ebook store!

Click to verify that you own this wallet and have control over it.

YiSin ebook (https://yisinnft.org/ebook) need to confirm whether you have the permission to read the e-book file.

This request will not trigger a blockchain transaction or cost any gas fees.`)

	hash := crypto.Keccak256Hash(data)
	signatureNoRecoverID := signatureByte[:len(signatureByte)-1]

	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	var book query.Book
	found := false

	for _, collName := range collections {
		coll := db.Collection(collName)
		filter := bson.M{"tokenId": id}
		err := coll.FindOne(context.TODO(), filter).Decode(&book)
		if err == nil {
			found = true
			break
		}
	}

	if !found {
		ctx.String(http.StatusBadRequest, "No book found with the given tokenId")
		return
	}

	address := book.Uploader
	publicKeyByte, err := hexutil.Decode(address)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	verified := crypto.VerifySignature(publicKeyByte, hash.Bytes(), signatureNoRecoverID)

	if !verified {
		ctx.String(http.StatusForbidden, "Signature verify fail")
		return
	}
	ctx.Next()
}
