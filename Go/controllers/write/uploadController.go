package write

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/http"
	"path"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-gonic/gin"
	ebook "yisinnft.org/m/v2/contract"
)

type UploadController struct {
	Instance *ebook.YiSinEBook
}

func (con UploadController) UploadEbook(ctx *gin.Context) {
	//	Get post params
	writer := ctx.PostForm("writer")
	publisher := ctx.PostForm("publisher")
	publishDate := ctx.PostForm("publishDate")
	isbn := ctx.PostForm("isbn")
	introduction := ctx.PostForm("introduction")
	chapter := ctx.PostForm("chapter")
	maxRentTimeStr := ctx.PostForm("maxRentTime")
	priceStr := ctx.PostForm("price")
	class := ctx.PostForm("class")
	amountStr := ctx.PostForm("amount")
	edition := ctx.PostForm("edition")
	pagesStr := ctx.PostForm("pages")
	uploader := ctx.PostForm("uploader")

	//	Transform amount(string) to amount(big.Int)
	amount := new(big.Int)
	var ok bool
	amount, ok = amount.SetString(amountStr, 10)
	if !ok {
		ctx.String(http.StatusBadRequest, "Amount transform fail")
		return
	}
	//	if amount > 100
	if amount.Cmp(big.NewInt(100)) == 1 {
		ctx.String(http.StatusBadRequest, "Exceed max supply 100")
		return
	}
	// Transform price(string) to price(float)
	price, err1 := strconv.ParseFloat(priceStr, 64)
	if err1 != nil {
		ctx.String(http.StatusBadRequest, "Price transform fail")
		return
	}
	//	Price from eth to wei
	ethPrice := big.NewFloat(price)
	weiValue := new(big.Int)
	ethPrice.Mul(ethPrice, big.NewFloat(params.Ether)).Int(weiValue)

	//	Transform pages(string) to pages(int)
	pages, err2 := strconv.Atoi(pagesStr)
	if err2 != nil {
		ctx.String(http.StatusBadRequest, "Pages transform fail")
		return
	}

	//	Transform maxRentTime(string) to maxRentTime(big.Int)
	maxRentTime := new(big.Int)
	maxRentTime, ok1 := maxRentTime.SetString(maxRentTimeStr, 10)
	if !ok1 {
		ctx.String(http.StatusBadRequest, "MaxRentTime transform fail")
		return
	}

	//	For contract send
	address := common.HexToAddress(uploader)
	_, _, _, _, _, _, _, _, _ = writer, publisher, publishDate, isbn, introduction, chapter, class, edition, pages
	file, err3 := ctx.FormFile("book")
	if err3 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"file_success": false,
		})
	}
	//	Verify file format
	extName := path.Ext(file.Filename)
	//	Set allow file extention
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
		".pdf":  true,
		".mp4":  true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		ctx.String(http.StatusBadRequest, "上傳文件類型不合法")
		return
	}

	dst := path.Join("./static/upload", file.Filename)
	ctx.SaveUploadedFile(file, dst)

	//	Call pravite functino to upload to blockchain.
	tx := con.uploadToBlockchain(amount, address, weiValue, maxRentTime)

	ctx.JSON(http.StatusOK, gin.H{
		"file_success":     true,
		"transaction_hash": tx,
	})
}

func (con UploadController) uploadToBlockchain(amount *big.Int, uploader common.Address, price *big.Int, time *big.Int) *types.Transaction {
	privateKey, err := crypto.HexToECDSA("24afe77abe16d1bf92de7e6b88590fda82d9fe20f3bd06582c935f7454b33002")
	if err != nil {
		log.Fatal(err)
	}
	//	Get public key from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//	Define ethereum node
	client, err1 := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err1 != nil {
		log.Fatal(err)
	}

	//	@param fromAddress company wallet address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//	Get nonce and gas pirce。
	nonce, err2 := client.PendingNonceAt(context.Background(), fromAddress)
	if err2 != nil {
		log.Fatal(err2)
	}
	gasPrice, err3 := client.SuggestGasPrice(context.Background())
	if err3 != nil {
		log.Fatal(err3)
	}

	//	Define trsaction auth
	chainID := big.NewInt(11155111)
	auth, err4 := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err4 != nil {
		log.Fatal(err4)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	tx, err5 := con.Instance.UploadEBook(auth, amount, uploader, price, time)
	if err5 != nil {
		log.Fatal(err5)
	}

	return tx
}
