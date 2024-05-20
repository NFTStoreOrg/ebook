package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/get"
)

//	路由的中間件：可以做權限檢查的函式，直接寫在要資料之前的函式
func GetNFTInfoAPIInit(r *gin.Engine) {
	getNFTInfoAPI := r.Group("/getnft")
	{
		//	可以傳入多個中間件
		getNFTInfoAPI.GET("/all", get.GetController{}.MiddleWare, get.GetController{}.GetAllNFT)
		getNFTInfoAPI.GET("/personal", get.GetController{}.MiddleWare, get.GetController{}.GetPersonalNFT)
		getNFTInfoAPI.GET("/remaining", get.GetController{}.MiddleWare, get.GetController{}.GetRemainingNFT)
		getNFTInfoAPI.POST("post")
	}
}
