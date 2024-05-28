package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/query"
)

func QueryNFTInit(r *gin.Engine) {
	instance := GetInstance()
	queryBookController := query.QueryBookController{Instance: instance}
	queryPersonalController := query.QueryPersonalController{Instance: instance}

	//	Configure router
	queryNFTApi := r.Group("/query")
	{
		queryBookApi := queryNFTApi.Group("/book")
		{
			queryBookApi.GET("/totalsupply", queryBookController.GetVarietyOfBook)
			queryBookApi.GET("/information/:id", queryBookController.GetBookInformation)
			queryBookApi.GET("/remain/:id", queryBookController.GetBookRemaining)
		}
		queryPersonalApi := queryNFTApi.Group("/:address")
		{
			queryPersonalApi.GET("/endtime/:id", queryPersonalController.GetRentBookEndTime)
			queryPersonalApi.GET("/rentedbook", queryPersonalController.GetPersonalRentedBook)
			queryPersonalApi.GET("/publish", queryPersonalController.GetPersonalPublish)
		}
	}
}
