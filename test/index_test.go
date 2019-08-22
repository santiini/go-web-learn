package test

import (
	initRouter "go-web/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

/// init 方法， 引入模块前执行的初始化函数
func init() {
	// 设置 mode=test, 解决 test 环境下 template 找不到的问题
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

// router("/")  测试
func TestIndexHtml(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method", "返回的html包含hello gin get method")
}

/*
	运行测试用例
	  1. 在项目根目录，执行 go test -v ./test
*/
