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

	/*
		路由分组
			1. 相同逻辑的代码集中处理
			2. 代码提供相同的路由前缀
	*/

	index := r.Group("/")
	{
		// 注册路由
		// gin.Context 集合了 request, Params 等的属性和方法
		// index.GET("/", retHelloGinAndMethod)
		// post
		// index.POST("/", retHelloGinAndMethod)
		// put
		// index.PUT("/", retHelloGinAndMethod)
		// delete
		// index.DELETE("/", retHelloGinAndMethod)
		// post
		// index.PATCH("/", retHelloGinAndMethod)
		// Head
		// index.HEAD("/", retHelloGinAndMethod)
		// Options
		// index.OPTIONS("/", retHelloGinAndMethod)

		// demo1: 路由的方法 Any 替代了所有的路由方法
		index.Any("", retHelloGinAndMethod)
	}

	userRoute := r.Group("/user")
	{
		// demo2 路由的参数
		userRoute.GET("/:name", handler.UserSave)
		userRoute.GET("", handler.UserSaveByQuery)
	}

	return r
}
