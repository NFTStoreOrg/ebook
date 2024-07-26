package query

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/models"
)

func SearchESDocument(ctx *gin.Context) {
	words := ctx.Param("title")
	resp, err := models.EsClient.Search().Index("children,other,reference,textbook,video").Query(&types.Query{
		MatchPhrase: map[string]types.MatchPhraseQuery{
			"title": {Query: words},
		},
	}).Do(context.Background())

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	var result []json.RawMessage
	for _, hit := range resp.Hits.Hits {
		result = append(result, hit.Source_)
	}
	ctx.JSON(http.StatusOK, result)
}
