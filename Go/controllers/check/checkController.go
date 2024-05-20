package check

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckController struct {
}

func (con CheckController) CheckAllNFT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "all-check",
	})
}
func (con CheckController) CheckPersonalNFT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "personal-check",
	})
}
func (con CheckController) CheckRemainingNFT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "remaining-check",
	})
}