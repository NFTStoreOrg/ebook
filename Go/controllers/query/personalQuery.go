package query

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	ebook "yisinnft.org/m/v2/contract"
)

type QueryPersonalController struct {
	Instance *ebook.YiSinEBook
}

func (con QueryPersonalController) GetRentBookEndTime(ctx *gin.Context) {
	idstr := ctx.Param("id")
	idBigInt, ok := new(big.Int).SetString(idstr, 10)
	//	Test id correct
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	addressStr := ctx.Param("address")
	address := common.HexToAddress(addressStr)
	index, err := con.Instance.RenterRentInfoIndex(nil, address, idBigInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Index_error": err})
	}

	info, err := con.Instance.RentInfos(nil, idBigInt, index)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"rent_info_index": err})
	}

	endTime := info.EndTime
	now := time.Now().Unix()
	duration := endTime.Sub(endTime, big.NewInt(now))
	durationToSecond := new(big.Int).Set(duration)
	secondInt64 := durationToSecond.Int64()

	days := secondInt64 / (24 * 60 * 60)
	hours := (secondInt64 % (24 * 60 * 60)) / (60 * 60)
	minutes := (secondInt64 % (60 * 60)) / 60
	seconds := secondInt64 % 60

	ctx.JSON(http.StatusOK, gin.H{
		"remaining_in_second": duration,
		"seconds_remaining":   seconds,
		"minutes_remaining":   minutes,
		"hours_remaining":     hours,
		"days_remaining":      days,
	})
}

func (con QueryPersonalController) GetPersonalRentedBook(ctx *gin.Context) {
	addressStr := ctx.Param("address")

	//	Use moralis api to get personal nft
	url := fmt.Sprintf("https://deep-index.moralis.io/api/v2.2/%s/nft?chain=sepolia&format=decimal&token_addresses%%5B0%%5D=0x62495223E379b2C752081d1dFd2D58C2B8E62Ec5&media_items=false", addressStr)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJub25jZSI6IjAxMzA5NzA1LTA0YjYtNDRkZS05YThlLTBjMzU5ZjA0ZGZhOCIsIm9yZ0lkIjoiMzg3NjI5IiwidXNlcklkIjoiMzk4Mjk2IiwidHlwZUlkIjoiNTcxMTgwZjYtYWUxZS00ZGU4LTk3NGQtYTJmMDljNDUwM2VlIiwidHlwZSI6IlBST0pFQ1QiLCJpYXQiOjE3MTI5MTI4MTYsImV4cCI6NDg2ODY3MjgxNn0.u6VKPRrHScpsHzf83hht5-E1UKKDz3DlQ6BMlayaFLg")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// Use json.Unmarshal, change json(body) to map(result)
	var result gin.H
	json.Unmarshal(body, &result)

	// extract token_id and metadata
	items := result["result"].([]interface{})
	var extractedItems []gin.H
	for _, item := range items {
		itemMap := item.(gin.H)
		tokenID := itemMap["token_id"].(string)
		metadata := itemMap["metadata"].(string)
		extractedItems = append(extractedItems, gin.H{
			"token_id": tokenID,
			"metadata": metadata,
		})
	}

	// Return message after extract
	ctx.JSON(http.StatusOK, extractedItems)
}

func (con QueryPersonalController) GetPersonalPublish(ctx *gin.Context) {
	address := common.HexToAddress(ctx.Param("address"))
	_ = address
	//	撈資料庫
}
