package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/routers" //	引用routers包調用方法
)

func main() {
	//	Initial API information
	r := gin.Default()

	//	設定信任代理ip
	trustedProxies := []string{"211.75.24.91"}
	err := r.SetTrustedProxies(trustedProxies)
	if err != nil {
		log.Fatal("Set trust proxies fail: ", err)
	}

	routers.QueryNFTInit(r)
	routers.UploadInit(r)

	r.Run(":8080") //	裡面可以寫端口
}
