package initrouter

import (
	"go-web/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}

// SetupRouter 注册路由
func SetupRouter() *gin.Engine {
	// engine 的生成和使用过程
	r := gin.Default()

	// demo1: 路由的方法
	// 注册路由
	// gin.Context 集合了 request, Params 等的属性和方法
	r.GET("/", retHelloGinAndMethod)

	// post
	r.POST("/", retHelloGinAndMethod)
	// put
	r.PUT("/", retHelloGinAndMethod)
	// delete
	r.DELETE("/", retHelloGinAndMethod)
	// post
	r.PATCH("/", retHelloGinAndMethod)
	// Head
	r.HEAD("/", retHelloGinAndMethod)
	// Options
	r.OPTIONS("/", retHelloGinAndMethod)

	// demo2 路由的参数
	r.GET("/user/:name", handler.UserSave)
	r.GET("/user", handler.UserSaveByQuery)

	return r
}
