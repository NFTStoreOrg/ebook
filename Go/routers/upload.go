package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/write"
)

func UploadInit(r *gin.Engine) {
	instance := GetInstance()

	uploadController := write.UploadController{Instance: instance}
	r.POST("/upload", uploadController.UploadEbook)
}
