package main

import (
	"log"

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
	//	Initial routers
	routers.QueryNFTInit(r)
	routers.UploadInit(r)

	r.Run(":8080")
}

