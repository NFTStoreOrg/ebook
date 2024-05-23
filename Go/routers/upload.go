package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/upload"
)

func UploadInit(r *gin.Engine) {
	uploadController := upload.UploadController{}
	r.POST("/upload", uploadController.UploadEbook)
}
