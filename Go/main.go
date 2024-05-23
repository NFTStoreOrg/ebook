package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"yisinnft.org/m/v2/routers" //	引用routers包調用方法
)

// `json: "title"`的用途為返回數據時變為小寫
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 使用form:"username"表示可以解析表單內容並且綁到這個結構內
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func main() {
	//	Initial API information
	r := gin.Default()

	//	設定信任代理ip
	trustedProxies := []string{"211.75.24.91"}
	err := r.SetTrustedProxies(trustedProxies)
	if err != nil {
		log.Fatal("Set trust proxies fail: ", err)
	}
	
	//	也可以直接返回結構
	r.GET("/json/struct", func(ctx *gin.Context) {
		a := &Article{
			Title:   "title",
			Desc:    "description",
			Content: "content",
		}
		ctx.JSON(200, a)
	})
	//	回傳jsonp，可以使用http://localhost:8080/jsonp/struct?callback=xxx
	//	將json文件傳入xxx callback function內，解決跨域問題（？
	r.GET("/jsonp/struct", func(ctx *gin.Context) {
		a := &Article{
			Title:   "title",
			Desc:    "description",
			Content: "content",
		}
		ctx.JSONP(200, a)
	})

	//	get請求傳值
	//	使用localhost:8080/?username=testuser&age=20傳
	r.GET("/getparam", func(ctx *gin.Context) {
		username := ctx.Query("username")
		age := ctx.Query("age")
		//	defaultQuery可以定義當沒有傳值時的預設數值
		page := ctx.DefaultQuery("page", "1")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	//	動態路由傳值 ！！比較重要
	//	用localhost:8080/user/20
	r.GET("/user/:uid", func(ctx *gin.Context) {
		uid := ctx.Param("uid")
		ctx.String(http.StatusOK, "userID=%s", uid)
	})

	//	post傳值
	//	使用表單post值
	r.POST("/doAddUser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		age := ctx.DefaultPostForm("age", "20")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	r.GET("/getUser/:username/:password", func(ctx *gin.Context) {
		user := &UserInfo{
			Username: ctx.Param("username"),
			Password: ctx.Param("password"),
		}
		if err := ctx.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			ctx.JSON(http.StatusOK, user)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	//	api路由分組：很多種功能的api可以分組開發，這樣表現比較清晰
	//	訪問時使用 /getnft/all, /getnft/personal, /getnft/remaining
	//	可以用這種方式分很多組
	routers.QueryNFTInit(r)

	r.Run(":8080") //	裡面可以寫端口
}
