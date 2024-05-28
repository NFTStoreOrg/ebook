package query

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	ebook "yisinnft.org/m/v2/contracts"
)

type QueryBookController struct {
	Instance *ebook.YiSinEBook
}

func (con QueryBookController) GetVarietyOfBook(ctx *gin.Context) {
	total, _ := con.Instance.TotalSupplyBook(nil)
	ctx.JSON(http.StatusOK, gin.H{
		"variety": total,
	})

}

func (con QueryBookController) GetBookInformation(ctx *gin.Context) {
	idstr := ctx.Param("id")
	idBigInt, ok := new(big.Int).SetString(idstr, 10)
	//	Test id correct
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	info, _ := con.Instance.BookInfos(nil, idBigInt)
	//	Return book information
	ctx.JSON(http.StatusOK, gin.H{
		"writer_address": info.Writer,
		"supply_amount":  info.SupplyAmount,
		"price":          info.RentPrice,
		"max_rent_time":  info.MaxRentTime,
	})
}

func (con QueryBookController) GetBookRemaining(ctx *gin.Context) {
	idstr := ctx.Param("id")
	idBigInt, ok := new(big.Int).SetString(idstr, 10)
	//	Test id correct
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	amount, _ := con.Instance.BooksOnRent(nil, idBigInt)
	info, _ := con.Instance.BookInfos(nil, idBigInt)
	total := info.SupplyAmount
	//	Calculate remaining nft.
	difference := new(big.Int).Sub(total, amount)

	ctx.JSON(http.StatusOK, gin.H{
		"remaining_amount": difference,
	})
}
