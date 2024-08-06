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
	mutex    sync.Mutex
)

func InitInstance() {
	mutex.Lock()
	once.Do(func() {
		client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
		if err != nil {
			log.Fatal(err)
		}
		address := common.HexToAddress("0xD658Ca5061B4e5bAbFAA49c70A52033dC1f98a78")
		inst, err := ebook.NewYiSinEBook(address, client)
		if err != nil {
			log.Fatal(err)
		}
		instance = inst
	})
	mutex.Unlock()
}

func GetInstance() *ebook.YiSinEBook {
	if instance == nil {
		InitInstance()
	}
	return instance
}
