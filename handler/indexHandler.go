package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Index 渲染 html 模板
func Index(context *gin.Context) {
	// gin.H{} 是传递到 index.tmpl 模板中的参数信息
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}
