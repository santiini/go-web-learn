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

	// static 静态文件的加载, 指定 /statics 路由映射到 ./statics 文件目录，可以映射图片等资源
	r.Static("/statics", "./statics")

	// 指定 .ico 图片
	r.StaticFile("/favicon.ico", "./../favicon.ico")

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
		userRoute.POST("/register", handler.UserRegister)
		userRoute.GET("/:name", handler.UserSave)
		userRoute.GET("", handler.UserSaveByQuery)
	}

	return r
}
