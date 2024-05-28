package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/write"
)

func UploadInit(r *gin.Engine) {
	instance := GetInstance()

	uploadController := write.UploadController{Instance: instance}
	rentController := write.RentController{Instance: instance}
	writeNFT := r.Group("/write")
	{
		writeNFT.POST("/upload", uploadController.UploadEbook)
		writeNFT.POST("/rent", rentController.RentBook)
	}

}
