package query

import (
	"context"
	"math/big"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	ebook "yisinnft.org/m/v2/contracts"
)

type QueryBookController struct {
	Instance *ebook.YiSinEBook
	DB       *mongo.Client
}

type Book struct {
	Title        string  `bson:"title,omitempty"`
	ISBN         string  `bson:"ISBN,omitempty"`
	Amount       int     `bson:"amount,omitempty"`
	Chapter      string  `bson:"chapter,omitempty"`
	Class        Class   `bson:"class,omitempty"`
	Edition      string  `bson:"edition,omitempty"`
	Introduction string  `bson:"introduction,omitempty"`
	Live         bool    `bson:"live,omitempty"`
	MaxRentTime  int     `bson:"maxRentTime,omitempty"`
	Pages        int     `bson:"pages,omitempty"`
	Price        float64 `bson:"price,omitempty"`
	PublishDate  string  `bson:"publishDate,omitempty"`
	Publisher    string  `bson:"publisher,omitempty"`
	Uploader     string  `bson:"uploader,omitempty"`
	Writer       string  `bson:"writer,omitempty"`
	TokenId      int     `bson:"tokenId,omitempty"`
}

type Class struct {
	ClassName string `bson:"class_name,omitempty"`
	Grade     int    `bson:"grade,omitempty"`
}

func (con QueryBookController) GetVarietyOfBook(ctx *gin.Context) {
	total, _ := con.Instance.TotalSupplyBook(nil)
	ctx.JSON(http.StatusOK, gin.H{
		"variety": total,
	})

}

func (con QueryBookController) GetBookInformation(ctx *gin.Context) {
	idstr := ctx.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Transform id fail",
		})
		return
	}

	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get collection fail",
		})
		return
	}
	var book Book
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

	if !found{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":"No book found with the given tokenId",
		})
		return
	}

	//	Return book information
	ctx.JSON(http.StatusOK, gin.H{
		"data":book,
	})
}

func (con QueryBookController) GetBookRemaining(ctx *gin.Context) {
	idstr := ctx.Param("id")
	idBigInt, ok := new(big.Int).SetString(idstr, 10)
	//	Test id correct
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	amount, _ := con.Instance.BooksOnRent(nil, idBigInt)
	info, _ := con.Instance.BookInfos(nil, idBigInt)
	total := info.SupplyAmount
	//	Calculate remaining nft.
	difference := new(big.Int).Sub(total, amount)

	ctx.JSON(http.StatusOK, gin.H{
		"remaining_amount": difference,
	})
}

func (con QueryBookController) GetClassOfBooks(ctx *gin.Context) {
	class := ctx.Param("class")
	//	Get collection
	collection := con.DB.Database("ebook").Collection(class)

	//	Attempt to find data
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var results []Book

	for cur.Next(context.Background()) {
		var result Book
		err := cur.Decode(&result)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failing while decode result",
			})
			return
		}

		results = append(results, result)
	}
	defer cur.Close(context.Background())

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})

}
