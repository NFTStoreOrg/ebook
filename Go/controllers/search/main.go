package search

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8/esapi"
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
	content := ctx.Param("title")

	//	Construct query request
	query := fmt.Sprintf(`{
		"query": {
			"multi_match": {
				"query": "%s",
				"fields": ["title", "writer"],
				"fuzziness": "AUTO"
			}
		}
	}`, content)

	size := 1000
	scores := true
	req := esapi.SearchRequest{
		Index:       []string{"children", "other", "reference", "textbook", "video"},
		Body:        strings.NewReader(query),
		Size:        &size,
		TrackScores: &scores,
	}

	res, err := req.Do(context.Background(), models.EsClient)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error getting response: "+err.Error())
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		ctx.String(http.StatusInternalServerError, "Error in response: "+res.String())
		return
	}

	var result []map[string]interface{}
	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		ctx.String(http.StatusInternalServerError, "Error parsing the response body: "+err.Error())
		return
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})

	for _, hit := range hits {
		var item map[string]interface{}
		source := hit.(map[string]interface{})["_source"]
		score := hit.(map[string]interface{})["_score"]

		item = source.(map[string]interface{})
		item["score"] = score

		result = append(result, item)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
