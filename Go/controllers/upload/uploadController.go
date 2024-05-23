package upload

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (con UploadController) UploadEbook(ctx *gin.Context) {
	writer := ctx.PostForm("writer")
	publisher := ctx.PostForm("publisher")
	publishDate := ctx.PostForm("publishDate")
	isbn := ctx.PostForm("isbn")
	introduction := ctx.PostForm("introduction")
	chapter := ctx.PostForm("chapter")
	maxRentTime := ctx.PostForm("maxRentTime")
	price := ctx.PostForm("price")
	class := ctx.PostForm("class")
	amount := ctx.PostForm("amount")
	edition := ctx.PostForm("edition")
	pages := ctx.PostForm("pages")
	_, _, _, _, _, _, _, _, _, _, _, _ = writer, publisher, publishDate, isbn, introduction, chapter, maxRentTime, price, class, amount, edition, pages
	file, err := ctx.FormFile("book")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"file_success": false,
		})
	}

	extName := path.Ext(file.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
		".pdf":  true,
		".mp4":  true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		ctx.String(http.StatusBadRequest, "上傳文件類型不合法")
		return
	}

	dst := path.Join("./static/upload", file.Filename)
	ctx.SaveUploadedFile(file, dst)
	ctx.JSON(http.StatusOK, gin.H{
		"file_success": true,
	})
}
