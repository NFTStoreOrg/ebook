package query

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"sync"

	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	ebook "yisinnft.org/m/v2/contracts"
)

type QueryBookController struct {
	Instance *ebook.YiSinEBook
	DB       *mongo.Client
}

type Book struct {
	Title        string  `bson:"title,omitempty" json:"title"`
	ISBN         string  `bson:"ISBN,omitempty" json:"ISBN"`
	Amount       int64   `bson:"amount,omitempty" json:"amount"`
	Chapter      string  `bson:"chapter,omitempty" json:"chapter"`
	Class        Class   `bson:"class,omitempty" json:"class"`
	Edition      string  `bson:"edition,omitempty" json:"edition"`
	Introduction string  `bson:"introduction,omitempty" json:"introduction"`
	Live         bool    `bson:"live,omitempty" json:"live"`
	MaxRentTime  int64   `bson:"maxRentTime,omitempty" json:"max_rent_time"`
	Pages        int     `bson:"pages,omitempty" json:"pages"`
	Price        float64 `bson:"price,omitempty" json:"price"`
	PublishDate  string  `bson:"publishDate,omitempty" json:"publish_date"`
	Publisher    string  `bson:"publisher,omitempty" json:"publisher"`
	Uploader     string  `bson:"uploader,omitempty" json:"uploader"`
	Writer       string  `bson:"writer,omitempty" json:"writer"`
	TokenId      int64   `bson:"tokenId,omitempty" json:"tokenId"`
	UploadTime   int64   `bson:"uploadTime,omitempty" json:"upload_time"`
	CoverImage   string  `bson:"coverImage" json:"cover_image"`
}

type Class struct {
	ClassName string `bson:"class_name,omitempty" json:"class_name"`
	Grade     int    `bson:"grade,omitempty" json:"grade"`
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

	if !found {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No book found with the given tokenId",
		})
		return
	}

	//	Return book information
	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
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
			fmt.Println(err.Error())
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

func (con QueryBookController) GetTextbookGrade(ctx *gin.Context) {
	gradeStr := ctx.Param("grade")

	grade, err := strconv.Atoi(gradeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	collection := con.DB.Database("ebook").Collection("textbook")

	filter := bson.M{"class.grade": grade}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	defer cur.Close(context.TODO())

	var results []Book

	for cur.Next(context.TODO()) {
		var result Book
		err := cur.Decode(&result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		results = append(results, result)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func (con QueryBookController) GetClassOfTwentyBooksForIndex(ctx *gin.Context) {
	class := ctx.Param("class")

	collection := con.DB.Database("ebook").Collection(class)

	//	Set find options to get newest 20 books information
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "uploadTime", Value: -1}})
	findOptions.SetLimit(20)

	var results []Book
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": "Error occur while searching data",
		})
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var result Book
		err := cur.Decode(&result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error while decoding data",
				"message": err.Error(),
			})
		}
		results = append(results, result)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

// Get latest 12 books for index, use merge sort
func (con QueryBookController) GetNewestTwelveBookForIndex(ctx *gin.Context) {
	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failing when search collections",
		})
		return
	}

	var wg sync.WaitGroup
	booksChannel := make(chan []Book)
	sortedBooksArray := make([][]Book, 0)

	//	Use a goroutine consuming books information in channel
	go func() {
		for book := range booksChannel {
			sortedBooksArray = append(sortedBooksArray, book)
		}
	}()

	for _, collName := range collections {
		wg.Add(1)

		//	Use mongoDB SetSort function to get sorted books array (by uploadTime desc)
		go func(collName string) {
			defer wg.Done()
			coll := db.Collection(collName)
			//	Find every collection's newest 12 books
			cur, _ := coll.Find(context.TODO(), bson.D{{}}, options.Find().SetSort(bson.D{{Key: "uploadTime", Value: -1}}).SetLimit(12))
			if cur == nil {
				return
			}
			defer cur.Close(context.TODO())

			//	Decode result into book, and append to results
			var results []Book
			for cur.Next(context.TODO()) {
				var book Book
				if err := cur.Decode(&book); err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error": "Error occur while decoding data",
					})
					return
				}
				results = append(results, book)
			}
			//	Put results into channel
			booksChannel <- results
		}(collName)
	}

	wg.Wait()
	close(booksChannel)

	result := mergeSortedArrays(sortedBooksArray)

	if len(result) > 12 {
		result = result[:12]
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func mergeSortedArrays(sortedArrays [][]Book) []Book {
	//	If only one array, return directly
	if len(sortedArrays) == 1 {
		return sortedArrays[0]
	}
	//	Merge every book arrays until length of array is 1
	for len(sortedArrays) > 1 {
		var newSortedArrays [][]Book
		for i := 0; i < len(sortedArrays); i += 2 {
			//	If this array is the last array, append to new arrays
			if i == len(sortedArrays)-1 {
				newSortedArrays = append(newSortedArrays, sortedArrays[i])
			} else { //	Merge two array into one sorted array, and append to new arrays
				merged := merge(sortedArrays[i], sortedArrays[i+1])
				newSortedArrays = append(newSortedArrays, merged)
			}
		}
		sortedArrays = newSortedArrays
	}

	return sortedArrays[0]
}

func merge(left, right []Book) []Book {
	var result []Book
	var i, j = 0, 0
	//	If left and right are not empty, compare the first one
	for i < len(left) && j < len(right) {
		if left[i].UploadTime > right[j].UploadTime {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func (con QueryBookController) GetLiveBook(ctx *gin.Context) {
	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failing when search collections",
		})
		return
	}

	var wg sync.WaitGroup
	booksChannel := make(chan Book)
	books := make([]Book, 0)

	go func() {
		for book := range booksChannel {
			books = append(books, book)
		}
	}()

	for _, collName := range collections {
		wg.Add(1)
		go func(collName string) {
			defer wg.Done()
			coll := db.Collection(collName)
			filter := bson.M{"live": true}
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
				booksChannel <- book
			}
		}(collName)
	}
	wg.Wait()
	close(booksChannel)

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
