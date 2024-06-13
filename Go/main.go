package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"POST", "GET", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
	}))
	
	//	Initial routers
	routers.QueryNFTInit(r)
	routers.UploadInit(r)

	r.Run(":8080")
}

