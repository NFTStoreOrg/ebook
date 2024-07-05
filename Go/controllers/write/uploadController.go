package write

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	ebook "yisinnft.org/m/v2/contracts"
)

type UploadController struct {
	Instance *ebook.YiSinEBook
	DB       *mongo.Client
}

// Process book information and return success or not.
func (con UploadController) UploadEbook(ctx *gin.Context) {
	//	Get post params
	title := ctx.PostForm("title")
	writer := ctx.PostForm("writer")
	publisher := ctx.PostForm("publisher")
	publishDate := ctx.PostForm("publishDate")
	uploader := ctx.PostForm("uploader")
	isbn := ctx.PostForm("isbn")
	introduction := ctx.PostForm("introduction")
	chapter := ctx.PostForm("chapter")
	maxRentTimeStr := ctx.PostForm("maxRentTime")
	priceStr := ctx.PostForm("price")
	className := ctx.PostForm("class")
	grade := ctx.PostForm("grade")
	amountStr := ctx.PostForm("amount")
	edition := ctx.PostForm("edition")
	pagesStr := ctx.PostForm("pages")
	liveStr := ctx.PostForm("live")

	//	Transform amount(string) to amount(big.Int)
	amount := new(big.Int)
	var ok bool
	amount, ok = amount.SetString(amountStr, 10)
	if !ok {
		ctx.String(http.StatusBadRequest, "Amount transform fail")
		return
	}
	//	if amount > 100
	if amount.Cmp(big.NewInt(1000)) == 1 {
		ctx.String(http.StatusBadRequest, "Amount exceed max supply 1000")
		return
	}

	isbnValid := checkISBNValid(isbn)
	if !isbnValid {
		ctx.String(http.StatusBadRequest, "Invalid isbn value")
		return
	}

	// Transform price(string) to price(float)
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	//	Price from eth to wei
	ethPrice := big.NewFloat(price)
	weiValue := new(big.Int)
	ethPrice.Mul(ethPrice, big.NewFloat(params.Ether)).Int(weiValue)

	//	Transform pages(string) to pages(int)
	pages, err := strconv.Atoi(pagesStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	//	Transform maxRentTime(string) to maxRentTime(big.Int)
	maxRentTime := new(big.Int)
	maxRentTime, ok1 := maxRentTime.SetString(maxRentTimeStr, 10)
	if !ok1 {
		ctx.String(http.StatusBadRequest, "MaxRentTime transform fail")
		return
	}

	live, err := strconv.ParseBool(liveStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	//	Address for contract send
	address := common.HexToAddress(uploader)

	file, err := ctx.FormFile("bookCover")
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		fmt.Println("bookCover")
		return
	}

	bookFile, err := ctx.FormFile("book")
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		fmt.Println("book")
		return
	}

	//	Verify file format
	extName := path.Ext(file.Filename)
	bookExtName := path.Ext(bookFile.Filename)
	//	Set allow file extention
	allowExtMap := map[string]bool{
		".jpg":  true,
		".JPG":  true,
		".png":  true,
		".jpeg": true,
		".JPEG": true,
		".PNG":  true,
		".tiff": true,
	}
	allowBookExtMap := map[string]bool{
		".pdf":  true,
		".PDF":  true,
		".epub": true,
		".txt":  true,
		".mp4":  true,
	}
	//	Check file format
	if _, ok := allowExtMap[extName]; !ok {
		ctx.String(http.StatusBadRequest, "Invalid file type (cover)")
		return
	}
	if _, ok := allowBookExtMap[bookExtName]; !ok {
		ctx.String(http.StatusBadRequest, "Invalid file type (book)")
		return
	}

	//	Get this book's tokenId to
	tokenId, err := con.Instance.TotalSupplyBook(nil)
	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
		return
	}
	tokenIdStr := tokenId.String()

	//	Generate file's path and name
	dst := path.Join("./static/upload", tokenIdStr+extName)
	bookDst := path.Join("./static/bookfile", tokenIdStr+bookExtName)

	//	Save file
	ctx.SaveUploadedFile(file, dst)
	ctx.SaveUploadedFile(bookFile, bookDst)

	httpDst := "https://yisinnft.org/images/" + tokenIdStr + extName

	tokenIdInt := tokenId.Int64()
	amountInt := amount.Int64()
	maxRentTimeInt := maxRentTime.Int64()
	gradeInt, err := strconv.Atoi(grade)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	//	Write metadata to database
	metaData := gin.H{
		"title":        title,
		"writer":       writer,
		"publisher":    publisher,
		"publishDate":  publishDate,
		"EISBN":        isbn,
		"introduction": introduction,
		"chapter":      chapter,
		"maxRentTime":  maxRentTimeInt,
		"price":        price,
		"class": gin.H{
			"class_name": className,
			"grade":      gradeInt,
		},
		"amount":     amountInt,
		"edition":    edition,
		"pages":      pages,
		"uploader":   uploader,
		"live":       live,
		"tokenId":    tokenIdInt,
		"uploadTime": time.Now().Unix(),
		"coverImage": httpDst,
	}
	collection := con.DB.Database("ebook").Collection(className)

	result, err := collection.InsertOne(context.Background(), metaData)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	tx, err := con.uploadToBlockchain(amount, address, weiValue, maxRentTime)
	if err != nil {
		ctx.String(http.StatusBadGateway, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"upload_status": true,
		"result":        result,
		"tx_hash":       tx.Hash().Hex(),
	})

}

func (con UploadController) uploadToBlockchain(amount *big.Int, uploader common.Address, price *big.Int, time *big.Int) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA("24afe77abe16d1bf92de7e6b88590fda82d9fe20f3bd06582c935f7454b33002")
	if err != nil {
		return nil, err
	}
	//	Get public key from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey, got %T", publicKey)
	}
	//	Define ethereum node
	client, err1 := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err1 != nil {
		return nil, err1
	}

	//	@param fromAddress company wallet address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//	Get nonce and gas pirce。
	nonce, err2 := client.PendingNonceAt(context.Background(), fromAddress)
	if err2 != nil {
		return nil, err2
	}
	gasPrice, err3 := client.SuggestGasPrice(context.Background())
	if err3 != nil {
		return nil, err3
	}

	//	Define trsaction auth
	chainID := big.NewInt(11155111)
	auth, err4 := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err4 != nil {
		return nil, err4
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	tx, err5 := con.Instance.UploadEBook(auth, amount, uploader, price, time)
	if err5 != nil {
		return nil, err5
	}

	return tx, nil
}

func checkISBNValid(isbn string) bool {
	//	Use ISBN-13
	if len(isbn) != 13 {
		return false
	}

	sum := 0
	lastCode := 0
	for i, digit := range isbn {
		if i == 12 {
			lastCode, _ = strconv.Atoi(string(digit))
		} else if i%2 == 0 {
			digit, _ := strconv.Atoi(string(digit))
			sum += digit
		} else {
			digit, _ := strconv.Atoi(string(digit))
			sum += 3 * digit
		}
	}

	valid := checkCode(lastCode, sum)
	return valid
}

func checkCode(lastCode int, sum int) bool {
	checkSum := sum % 10
	if checkSum == 0 {
		return lastCode == checkSum
	}

	code := 10 - checkSum
	return lastCode == code
}
