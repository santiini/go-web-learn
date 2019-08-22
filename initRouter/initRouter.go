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
	r.GET("/ping", func(c *gin.Context) {

		// 原理上，利用 w.writer 生成 json 返回结果
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/string", stringHandler)

	r.GET("/struct", strucHandler)

	return r
}

// Result 统一的返回结果
type Result struct {
	Status string
	code   int64
	data   interface{}
}

// map 作为 JSON 的参数, 返回 JSON
func strucHandler(c *gin.Context) {
	// 方式1
	// c.JSON(200, gin.H{
	// 	"status": "success",
	// 	"code":   200,
	// 	"data":   []string{"first", "second"},
	// })

	res := map[string]interface{}{
		"status": "success",
		"code":   200,
		"data":   []string{"first", "second"},
	}

	// 方式2
	c.JSON(200, res)

	// jsonRes := Result{"success", 200, []string{"first", "second", "third"}}

	// 方式3测试, 不能用  struct
	// c.JSON(200, jsonRes)
}

// 返回 string
func stringHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello gin")
}
