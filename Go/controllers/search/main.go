package search

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"yisinnft.org/m/v2/models"
)

type SearchController struct {
	DB *mongo.Client
}

func (con SearchController) CreateIndex(ctx *gin.Context) {
	title := ctx.Param("collection")
	models.CreateESIndex(title)
	ctx.JSON(http.StatusOK, gin.H{"status": true})
}

func (con SearchController) SyncDocument(ctx *gin.Context) {
	db := con.DB.Database("ebook")
	collections, err := db.ListCollectionNames(context.Background(), bson.M{})

	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
	}

	//	Find all database data
	var wg sync.WaitGroup
	bookChannel := make(chan models.Book, 5)
	books := make([]models.Book, 0)

	go func() {
		for book := range bookChannel {
			books = append(books, book)
		}
	}()

	for _, collection := range collections {
		wg.Add(1)

		go func(collName string) {
			defer wg.Done()

			coll := db.Collection(collName)
			cur, _ := coll.Find(context.Background(), bson.D{})

			if cur == nil {
				return
			}

			for cur.Next(context.Background()) {
				var result models.Book

				err = cur.Decode(&result)
				if err != nil {
					ctx.String(http.StatusInternalServerError, err.Error())
					return
				}
				bookChannel <- result
			}

		}(collection)
	}
	wg.Wait()
	close(bookChannel)

	//	add to elasticsearch
	for _, value := range books {
		_, err := models.CreateESDocument(value)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func (con SearchController) FuzzySearch(ctx *gin.Context) {
	title := ctx.Param("title")

	//	Construct query request
	query := types.Query{
		Match: map[string]types.MatchQuery{
			"title": {Query: title, Fuzziness: "AUTO"},
		},
	}

	//	Execute query
	res, err := models.EsClient.Search().Index("children,other,reference,textbook,video").
		Query(&query).
		Size(1000).
		TrackScores(true). //	Enable scores
		Do(context.Background())

	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
		return
	}

	var result []map[string]interface{}

	for _, hit := range res.Hits.Hits {
		var item map[string]interface{}
		//	Unmarshal json to item struct
		if err = json.Unmarshal(hit.Source_, &item); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		//	Add score key
		item["score"] = hit.Score_
		result = append(result, item)
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
