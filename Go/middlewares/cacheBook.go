package middlewares

import "github.com/gin-gonic/gin"

type CacheBookMiddleware struct{}

func (con CacheBookMiddleware) BookInfoMiddleware(ctx *gin.Context) {
	id := ctx.Param("id")
	cacheKey := "book_info_" + id

	getFromRedis(ctx, cacheKey)
}

func (con CacheBookMiddleware) ClassOfBookMiddleware(ctx *gin.Context) {
	class := ctx.Param("class")
	cacheKey := "book_" + class

	getFromRedis(ctx, cacheKey)
}

func (con CacheBookMiddleware) TextbookGradeMiddleware(ctx *gin.Context) {
	gradeStr := ctx.Param("grade")
	cacheKey := "textbook_grade" + gradeStr

	getFromRedis(ctx, cacheKey)
}
