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
	queryNFTApi := r.Group("/query")
	{
		queryBookApi := queryNFTApi.Group("/book")
		{
			queryBookApi.GET("/totalsupply", queryBookController.GetVarietyOfBook)
			queryBookApi.GET("/information/:id", queryBookController.GetBookInformation)
			queryBookApi.GET("/remain/:id", queryBookController.GetBookRemaining)
			queryBookApi.GET("/:class",queryBookController.GetClassOfBooks)
		}
		queryPersonalApi := queryNFTApi.Group("/:address")
		{
			queryPersonalApi.GET("/endtime/:id", queryPersonalController.GetRentBookEndTime)
			queryPersonalApi.GET("/rentedbook", queryPersonalController.GetPersonalRentedBook)
			queryPersonalApi.GET("/publish", queryPersonalController.GetPersonalPublish)
		}
	}
}
