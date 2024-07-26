package write

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yisinnft.org/m/v2/controllers/query"
	"yisinnft.org/m/v2/models"
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

	//	Update elasticsearch info
	var book models.Book
	if err := mapstructure.Decode(updateData, &book); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	err := models.UpdateESDocument(book)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	//	Update DB info
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

	//	Find bookId's uploader.
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
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
		ctx.String(http.StatusNotFound, "No book found with the given tokenId")
	}

	address := book.Uploader

	//	Change public key to address
	recoveredAddress := crypto.PubkeyToAddress(*recoveredPubKey).Hex()
	lowerCaseAddress := strings.ToLower(recoveredAddress)

	if ok := address == lowerCaseAddress; !ok {
		ctx.String(http.StatusForbidden, "Signature not match.")
	}

	ctx.Next()
}
