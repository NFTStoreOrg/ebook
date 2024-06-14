package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/query"
)

func QueryNFTInit(r *gin.Engine) {
	instance := GetInstance()
	db := GetDB()
	queryBookController := query.QueryBookController{Instance: instance, DB: db}
	queryPersonalController := query.QueryPersonalController{Instance: instance, DB: db}

	//	Configure router
	queryBookApi := r.Group("/book")
	{
		queryBookApi.GET("/totalsupply", queryBookController.GetVarietyOfBook)
		queryBookApi.GET("/:id", queryBookController.GetBookInformation)
		queryBookApi.GET("/remain/:id", queryBookController.GetBookRemaining)
		queryBookApi.GET("/page/:class", queryBookController.GetClassOfBooks)
		queryBookApi.GET("/live", queryBookController.GetLiveBook)
		queryBookApi.GET("/index/:class", queryBookController.GetClassOfTwentyBooksForIndex)
		queryBookApi.GET("/index", queryBookController.GetNewestTwelveBookForIndex)
	}
	queryPersonalApi := r.Group("/:address")
	{
		queryPersonalApi.GET("/endtime/:id", queryPersonalController.GetRentBookEndTime)
		queryPersonalApi.GET("/rentedbook", queryPersonalController.GetPersonalRentedBook)
		queryPersonalApi.GET("/publish", queryPersonalController.GetPersonalPublish)
		queryPersonalApi.GET("/:id/read", queryPersonalController.AddressHaveRentedBook)
		queryPersonalApi.GET("/:id/:signature", queryPersonalController.VerifySignatureMiddleWare, queryPersonalController.CheckPermissionToAccessFileMiddleWare, queryPersonalController.GetBookFile)
	}
}
