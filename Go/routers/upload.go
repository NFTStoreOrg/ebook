package routers

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	ebook "yisinnft.org/m/v2/contract"
	"yisinnft.org/m/v2/controllers/write"
)

func UploadInit(r *gin.Engine) {

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
	uploadController := write.UploadController{Instance: instance}
	r.POST("/upload", uploadController.UploadEbook)
}
