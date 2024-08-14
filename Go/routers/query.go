package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/query"
	"yisinnft.org/m/v2/middlewares"
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
		queryBookApi.GET("/:id", middlewares.CacheBookMiddleware{}.BookInfoMiddleware, queryBookController.GetBookInformation)
		queryBookApi.GET("/remain/:id", queryBookController.GetBookRemaining)
		queryBookApi.GET("/page/:class", middlewares.CacheBookMiddleware{}.ClassOfBookMiddleware, queryBookController.GetClassOfBooks)
		queryBookApi.GET("/page/textbook/:grade", middlewares.CacheBookMiddleware{}.TextbookGradeMiddleware, queryBookController.GetTextbookGrade)
		queryBookApi.GET("/live", middlewares.CacheMiddleware{}.LiveBookMiddleware, queryBookController.GetLiveBook)
		queryBookApi.GET("/index/:class", middlewares.CacheMiddleware{}.ClassBookMiddleware, queryBookController.GetClassOfTwentyBooksForIndex)
		queryBookApi.GET("/index", middlewares.CacheMiddleware{}.NewestBookMiddleware, queryBookController.GetNewestTwelveBookForIndex)
	}
	queryPersonalApi := r.Group("/:address")
	{
		queryPersonalApi.GET("/endtime/:id", queryPersonalController.GetRentBookEndTime)
		queryPersonalApi.GET("/rentedbook", middlewares.CachePersonalMiddleware{}.PersonalRentedMiddleware, queryPersonalController.GetPersonalRentedBook)
		queryPersonalApi.GET("/publish", middlewares.CachePersonalMiddleware{}.PersonalPublishMiddleware, queryPersonalController.GetPersonalPublish)
		queryPersonalApi.GET("/:id/read", queryPersonalController.AddressHaveRentedBook)
		queryPersonalApi.GET("/:id/:signature", queryPersonalController.VerifySignatureMiddleWare, queryPersonalController.CheckPermissionToAccessFileMiddleWare, queryPersonalController.GetBookFile)
	}
}
