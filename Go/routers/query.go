package routers

import (
	"log"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	ebook "yisinnft.org/m/v2/contract"
	"yisinnft.org/m/v2/controllers/query"
)

func QueryNFTInit(r *gin.Engine) {
	//	Initial ethereum node
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}
	//	Initial contract
	address := common.HexToAddress("0x790e48C4F57F4415b9Aed58157A6A8436ea094A6")
	//	Instance is contract
	instance, err := ebook.NewYiSinEBook(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
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
