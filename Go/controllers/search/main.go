package search

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"yisinnft.org/m/v2/models"
)

type SearchController struct {
	DB *mongo.Client
}

func (con SearchController) CreateIndex(ctx *gin.Context) {
	title := ctx.Param("collection")
	models.CreateESIndex(title)
}

func (con SearchController) SyncDocument(ctx *gin.Context) {
	
}
