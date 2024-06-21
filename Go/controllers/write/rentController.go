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
	bookIdStr := ctx.Param("bookId")
	timeStr := ctx.PostForm("time")
	renter := ctx.Param("address") //	未來要傳給api的參數，告訴他是誰要借書
	_ = renter

	bookId := new(big.Int)
	var ok bool
	bookId, ok = bookId.SetString(bookIdStr, 10)
	if !ok {
		ctx.String(http.StatusBadRequest, "BookId transform fail.")
		return
	}
	info, _ := con.Instance.BookInfos(nil, bookId)
	price := info.RentPrice
	maxTime := info.MaxRentTime

	time := new(big.Int)
	time, ok1 := time.SetString(timeStr, 10)
	if !ok1 {
		ctx.String(http.StatusBadRequest, "Transform time fail")
		return
	}
	if time.Cmp(maxTime) == 1 {
		ctx.String(http.StatusBadRequest, "Time exceed max rent time")

		return
	}
	//	這裡開始將來要改成call api取得私鑰
	apiPrivateKey := ""
	privateKey, err := crypto.HexToECDSA(apiPrivateKey)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	auth := getTransactionAuth(privateKey, price, ctx)

	tx, err := con.Instance.RentBook(auth, bookId, time)
	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
		return
	}
	txHash := tx.Hash().Hex()
	ctx.JSON(http.StatusOK, gin.H{
		"rent_status": true,
		"tx_hash":     txHash,
	})
}

func (con RentController) ReturnBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("id")
	addressStr := ctx.PostForm("address")

	_ = addressStr

	bookIdBig := new(big.Int)

	bookIdBig, ok := bookIdBig.SetString(bookIdStr, 10)

	if !ok {
		ctx.String(http.StatusBadRequest, "Transform bookId fail")
		return
	}

	apiPrivateKey := "" // Call api to get key.

	privateKey, err := crypto.HexToECDSA(apiPrivateKey)

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	auth := getTransactionAuth(privateKey, big.NewInt(0), ctx)
	if auth == nil {
		return
	}

	tx, err := con.Instance.ReturnBook(auth, bookIdBig)
	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
		return
	}

	txHash := tx.Hash().Hex()
	ctx.JSON(http.StatusOK, gin.H{
		"return_status": true,
		"tx_hash":       txHash,
	})
}

func getTransactionAuth(privateKey *ecdsa.PrivateKey, price *big.Int, ctx *gin.Context) *bind.TransactOpts {
	//	Get node client
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return nil
	}

	//	Get public key and public key ecdsa
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		ctx.String(http.StatusInternalServerError, "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//	Get nonce and gas price.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return nil
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return nil
	}

	//	Set chain id and new an auth
	chainID := big.NewInt(11155111)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return nil
	}

	//	Set auth information
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = price
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth
}
