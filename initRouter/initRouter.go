package initrouter

import (
	"go-web/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter 注册路由
func SetupRouter() *gin.Engine {
	// engine 的生成和使用过程
	r := gin.Default()

	// 模板渲染: LoadHTMLGlob 使用 templates/ 下的模板
	// 根据 mode 加载不同的路径
	if mode := gin.Mode(); mode == gin.TestMode {
		r.LoadHTMLGlob("./../templates/*")
	} else {
		r.LoadHTMLGlob("templates/*")
	}

	/*
		路由分组
			1. 相同逻辑的代码集中处理
			2. 代码提供相同的路由前缀
	*/

	index := r.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRoute := r.Group("/user")
	{
		// demo2 路由的参数
		userRoute.GET("/:name", handler.UserSave)
		userRoute.GET("", handler.UserSaveByQuery)
	}

	return r
}
