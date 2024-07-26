package routers

import (
	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/controllers/query"
)

func SearchInit(r *gin.Engine) {
	r.GET("/es/:title",query.SearchESDocument)
}
