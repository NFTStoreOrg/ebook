package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	`json: "title"`的用途為返回數據時變為小寫
type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	//	可get post put delete
	r.GET("/", func(c *gin.Context) {
		//	http.StatusOK = 200
		c.String(http.StatusOK, "值：%v", "hello gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})
	//	可以返回手寫json
	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H {
			"success": true,
			"msg": "test message",
		})
	})
	//	也可以直接返回結構
	r.GET("/json/struct", func(ctx *gin.Context) {
		a := &Article{
			Title: "title",
			Desc: "description",
			Content: "content",
		}
		ctx.JSON(200, a)
	})
	//	回傳jsonp，可以使用http://localhost:8080/jsonp/struct?callback=xxx
	//	將json文件傳入xxx callback function內，解決跨域問題（？
	r.GET("/jsonp/struct", func(ctx *gin.Context) {
		a := &Article{
			Title: "title",
			Desc: "description",
			Content: "content",
		}
		ctx.JSONP(200, a)
	})
	r.Run(":8080") //	裡面可以寫端口
}
