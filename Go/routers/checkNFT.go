package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/check"
)

//	接下來，利用controller 來拆分所有路由的業務邏輯，將func移動到controllers內
func CheckNFTInfoInit(r *gin.Engine) {
	checkNFTInfoAPI := r.Group("/checknft")
	{
		checkNFTInfoAPI.GET("/all", check.CheckController{}.CheckAllNFT)
		checkNFTInfoAPI.GET("/personal", check.CheckController{}.CheckPersonalNFT)
		checkNFTInfoAPI.GET("/remaining", check.CheckController{}.CheckRemainingNFT)
	}
}