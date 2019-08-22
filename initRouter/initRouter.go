package initrouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter 注册路由
func SetupRouter() *gin.Engine {
	// engine 的生成和使用过程
	r := gin.Default()

	// 注册路由
	// gin.Context 集合了 request, Params 等的属性和方法
	r.GET("/", func(c *gin.Context) {
		// 原理上，利用 w.writer 生成 json 返回结果
		c.String(http.StatusOK, "hello gin")
	})

	// post
	r.POST("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin post method")
	})
	// put
	r.PUT("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin put method")
	})
	// delete
	r.DELETE("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin delete method")
	})
	// post
	r.PATCH("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin patch method")
	})
	// Head
	r.HEAD("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin head method")
	})
	// Options
	r.OPTIONS("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin options method")
	})

	return r
}
