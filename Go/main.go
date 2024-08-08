package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/models"
	"yisinnft.org/m/v2/routers"
)

func main() {
	//	Initial API information
	r := gin.Default()

	//	Set trust poxy ip
	trustedProxies := []string{"211.75.24.91"}
	err := r.SetTrustedProxies(trustedProxies)
	if err != nil {
		log.Fatal("Set trust proxies fail: ", err)
	}

	//	Set CORS policy
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "PATCH", "PUT"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//	Initial tokenId in redis
	tokenId, _ := routers.GetInstance().TotalSupplyBook(nil)
	tokenId64 := tokenId.Int64()
	err = models.RedisClient.SetNX(context.Background(), "tokenId", tokenId64, 0).Err()
	if err != nil {
		log.Fatal("Error setting redis key:", err)
	}

	//	Initial routers
	routers.QueryNFTInit(r)
	routers.UploadInit(r)
	routers.SearchInit(r)

	r.Run(":8080")
}
