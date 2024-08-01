package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/write"
	"yisinnft.org/m/v2/controllers/query"
)

func UploadInit(r *gin.Engine) {
	instance := GetInstance()
	db := GetDB()

	uploadController := write.UploadController{Instance: instance, DB: db}
	rentController := write.RentController{Instance: instance}
	modifyController := write.ModifyController{DB: db}
	queryPersonalController := query.QueryPersonalController{Instance: instance, DB: db}

	r.POST("/book/:address/:signature", queryPersonalController.VerifySignatureMiddleWare, uploadController.UploadEbook)
	r.PUT("/book/:id/:address", rentController.RentBook)
	r.DELETE("/book/:id/:address", rentController.ReturnBook)
	r.PATCH("/book/:id/:signature", modifyController.VerifySignatureMiddleWare, modifyController.AdjustBookInformation)
}