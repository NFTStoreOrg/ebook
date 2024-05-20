package get

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

// 這裡演示繼承controller，其實就是繼承結構
type GetController struct {
	BaseController
}

// 將方法加入結構(con GetController)，可以用來繼承
func (con GetController) GetAllNFT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "all",
	})
}

func (con GetController) GetPersonalNFT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "personal",
	})
}

func (con GetController) GetRemainingNFT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "remaining",
	})
}
func (con GetController) PostNFTInfo(ctx *gin.Context) {
	//	從這裡接收post過來的資料，PostForm裡面的參數是前端input定義的name
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.PostForm("age")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"passward": password,
		"age":      age,
	})
}
func (con GetController) TestSuccess(ctx *gin.Context) {
	con.Success(ctx)
}
func (con GetController) TestError(ctx *gin.Context) {
	con.Error(ctx)
}

func (con GetController) MiddleWare(ctx *gin.Context) {
	start := time.Now().UnixNano()

	fmt.Println("Middleware be run")
	//	當調用這個Next()函式時，會先執行中間件之後的函式，執行完再執行下面的東西
	ctx.Next()
	fmt.Println("Middleware2 be run")

	end := time.Now().UnixNano()
	fmt.Println(end - start)

	//	在中間件中使用goroutine必須要將ctx複製一份才可以
	cCp := ctx.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine be runed" + cCp.Request.URL.Path)
	}() //	這是goroutine的寫法，在執行時不影響用戶讀取資料
}
