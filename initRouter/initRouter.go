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

	return r
}
