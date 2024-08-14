package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"yisinnft.org/m/v2/models"
)

type CacheMiddleware struct{}

func getFromRedis(ctx *gin.Context, cacheKey string) {
	data, err := models.RedisClient.Get(context.Background(), cacheKey).Result()

	if err == redis.Nil {
		return
	} else if err != nil {
		log.Fatal(err.Error())
		return
	}

	var result []models.Book
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
	ctx.Abort()
}

func (con CacheMiddleware) ClassBookMiddleware(ctx *gin.Context) {
	class := ctx.Param("class")

	cacheKey := "index_" + class
	getFromRedis(ctx, cacheKey)
}

func (con CacheMiddleware) NewestBookMiddleware(ctx *gin.Context) {
	cacheKey := "index_newest"
	getFromRedis(ctx, cacheKey)
}

func (con CacheMiddleware) LiveBookMiddleware(ctx *gin.Context) {
	cacheKey := "index_live"
	getFromRedis(ctx, cacheKey)
}
