package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSave 保存用户, router("/user/:name") get
func UserSave(context *gin.Context) {
	username := context.Param("name")

	context.String(http.StatusOK, "用户"+username+"已经保存")
}

// UserSaveByQuery router("/user?name=lisi&age=18")
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	// age := context.Query("age")

	// context.DefaultQuery 方法，在获取时，如果没有该值则赋给一个默认值。
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}
