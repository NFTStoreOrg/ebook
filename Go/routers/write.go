package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/write"
)

func UploadInit(r *gin.Engine) {
	instance := GetInstance()
	db := GetDB()

	uploadController := write.UploadController{Instance: instance, DB: db}
	rentController := write.RentController{Instance: instance}

	r.POST("/book", uploadController.UploadEbook)
	r.POST("/book/:id/:address", rentController.RentBook)
	r.DELETE("/book/:id/:address", rentController.ReturnBook)
}
