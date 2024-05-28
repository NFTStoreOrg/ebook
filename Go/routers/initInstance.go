package routers

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ebook "yisinnft.org/m/v2/contracts"
)

var (
	instance *ebook.YiSinEBook
	once     sync.Once
)

func InitInstance() {
	once.Do(func() {
		client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
		if err != nil {
			log.Fatal(err)
		}
		address := common.HexToAddress("0x790e48C4F57F4415b9Aed58157A6A8436ea094A6")
		inst, err := ebook.NewYiSinEBook(address, client)
		if err != nil {
			log.Fatal(err)
		}
		instance = inst
	})
}

func GetInstance() *ebook.YiSinEBook {
	if instance == nil {
		InitInstance()
	}
	return instance
}
