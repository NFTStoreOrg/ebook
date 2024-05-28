package write

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gin-gonic/gin"
	ebook "yisinnft.org/m/v2/contracts"
)

type UploadController struct {
	Instance *ebook.YiSinEBook
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
	live := ctx.PostForm("live")

	//	Transform amount(string) to amount(big.Int)
	amount := new(big.Int)
	var ok bool
	amount, ok = amount.SetString(amountStr, 10)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"information_fail": "Amount transform fail",
		})
		return
	}
	//	if amount > 100
	if amount.Cmp(big.NewInt(1000)) == 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"information_fail": "Exceed max supply 100",
		})
		return
	}
	// Transform price(string) to price(float)
	price, err1 := strconv.ParseFloat(priceStr, 64)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"information_fail": "Price transform fail",
		})
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
	file, err3 := ctx.FormFile("book")
	if err3 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"file_success": false,
		})
		return
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
	//	Get this book's tokenId to
	tokenId, err6 := con.Instance.TotalSupplyBook(nil)
	if err6 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"token_id_success": false,
		})
		return
	}
	tokenId = tokenId.Sub(tokenId, big.NewInt(1))
	tokenIdStr := tokenId.String()
	//	Generate file's path and name
	dst := path.Join("./static/upload", className, tokenIdStr+extName)
	ctx.SaveUploadedFile(file, dst)

	//	Using goroutine to asynchronously execute blockchain upload books
	//	Struct for blockchain result
	type Result struct {
		Message *types.Transaction
		Error   error
	}
	//	Channel to receive the result
	ch := make(chan Result)
	//	Call pravite functino to upload to blockchain.
	go func() {
		tx, err := con.uploadToBlockchain(amount, address, weiValue, maxRentTime)

		if tx == nil {
			ch <- Result{nil, err}
			return
		}
		ch <- Result{tx, nil}
	}()

	//	Write metadata
	metaData := gin.H{
		"title": title,
		"type":  "Object",
		"properties": gin.H{
			"writer": gin.H{
				"type":        "string",
				"description": writer,
			},
			"publisher": gin.H{
				"type":        "string",
				"description": publisher,
			},
			"publishDate": gin.H{
				"type":        "string",
				"description": publishDate,
			},
			"ISBN": gin.H{
				"type":        "string",
				"description": isbn,
			},
			"introduction": gin.H{
				"type":        "string",
				"description": introduction,
			},
			"chapter": gin.H{
				"type":        "string",
				"description": chapter,
			},
			"maxRentTime": gin.H{
				"type":        "string",
				"description": maxRentTime,
			},
			"price": gin.H{
				"type":        "string",
				"description": price,
			},
			"class": gin.H{
				"type": "object",
				"description": gin.H{
					"class_name": className,
					"grade":      grade,
				},
			},
			"amount": gin.H{
				"type":        "string",
				"description": amount,
			},
			"edition": gin.H{
				"type":        "string",
				"description": edition,
			},
			"pages": gin.H{
				"type":        "string",
				"description": pages,
			},
			"uploader": gin.H{
				"type":        "string",
				"description": uploader,
			},
			"live": gin.H{
				"type":        "string",
				"descriptino": live,
			},
		},
	}

	fileData, err := json.MarshalIndent(metaData, "", "    ")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot generate JSON",
		})
		return
	}
	if grade != "0" {
		filePath := path.Join("./metadata", className, grade, tokenIdStr+".json")
		err = os.WriteFile(filePath, fileData, 0644)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Run failed while writing metadata file",
			})
			return
		}
	} else {
		filePath := path.Join("./metadata", className, tokenIdStr+".json")
		err = os.WriteFile(filePath, fileData, 0644)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Run failed while writing metadata file",
			})
			return
		}
	}

	res := <-ch
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"blockchain_error": res.Error,
		})
	} else {

		tx := res.Message
		ctx.JSON(http.StatusOK, gin.H{
			"file_success":     true,
			"transaction_hash": tx.Hash().Hex(),
		})
	}

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
