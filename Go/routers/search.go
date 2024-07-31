package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/search"
)

func SearchInit(r *gin.Engine) {
	db := GetDB()

	r.GET("/es/:title",search.SearchController{}.FuzzySearch)
	r.POST("/:collection", search.SearchController{}.CreateIndex)
	r.GET("/sync/es", search.SearchController{DB: db}.SyncDocument)
}
