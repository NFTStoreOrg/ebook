package get

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {}

func (con BaseController) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
func (con BaseController) Error(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "error",
	})
}