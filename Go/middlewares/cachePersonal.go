package middlewares

import "github.com/gin-gonic/gin"

type CachePersonalMiddleware struct{}

func (con CachePersonalMiddleware) PersonalRentedMiddleware(ctx *gin.Context) {
	address := ctx.Param("address")
	cacheKey := "personal_rented_" + address

	getFromRedis(ctx, cacheKey)
}

func(con CachePersonalMiddleware) PersonalPublishMiddleware(ctx *gin.Context){
	address := ctx.Param("address")
	cacheKey := "personal_publish_" + address

	getFromRedis(ctx, cacheKey)
}