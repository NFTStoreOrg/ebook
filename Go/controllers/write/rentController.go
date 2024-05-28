package write

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	ebook "yisinnft.org/m/v2/contracts"
)

type RentController struct {
	Instance *ebook.YiSinEBook
}

func (con RentController) RentBook(ctx *gin.Context) {
	bookIdStr := ctx.PostForm("bookId")
	timeStr := ctx.PostForm("time")

	bookId := new(big.Int)
	var ok bool
	bookId, ok = bookId.SetString(bookIdStr, 10)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"bookId_fail": "BookId transform fail.",
		})
	}
	info, _ := con.Instance.BookInfos(nil, bookId)
	price := info.RentPrice

	time := new(big.Int)
	time, ok1 := time.SetString(timeStr, 10)
	if !ok1 {
		ctx.String(http.StatusBadRequest, "Time transform fail")
		return
	}
	//	這裡開始將來要改成call api取得私鑰
	apiPrivateKey := ""
	privateKey, err := crypto.HexToECDSA(apiPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get private key fail",
		})
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot assert type: publicKey is not of type *ecdsa.PublicKey",
		})
		return
	}

	client, err1 := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Get ETH client fail",
		})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err2 := client.PendingNonceAt(context.Background(), fromAddress)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get nonce fail",
		})
		return
	}

	gasPrice, err3 := client.SuggestGasPrice(context.Background())
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get gas price fail",
		})
		return
	}

	chainID := big.NewInt(11155111)
	auth, err4 := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err4 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Get transaction auth fail",
		})
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = price
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	tx, err5 := con.Instance.RentBook(auth, bookId, time)
	if err5 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err5,
			"message": "Failed to rent a book on the blockchain",
		})
		return
	}
	txHash := tx.Hash().Hex()
	ctx.JSON(http.StatusOK, gin.H{
		"rent_status": true,
		"tx_hash":     txHash,
	})

}
