package test

import (
	initRouter "go-web/initRouter"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// router("/user/:name") get
func TestUserSave(t *testing.T) {
	username := "xiaotao"
	router := initRouter.SetupRouter()

	// 生成一条 httptest 的请求记录，NewRecorder() 是 http.ResponseWriter 的实现
	w := httptest.NewRecorder()

	// 根据参数生成请求，各种配置、参数在这里
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)

	// 根据路由 router 构建 handleFunc(w http.ResponseWriter, req *http.Request), 起服务
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户"+username+"已经保存", w.Body.String())
}

// TestUserSave router("/user?name=xiaotaot&age=18") get
func TestUserSaveQuery(t *testing.T) {
	username := "xiaotaot"
	age := 18
	router := initRouter.SetupRouter()

	// 生成一条 httptest 的请求记录，NewRecorder() 是 http.ResponseWriter 的实现
	w := httptest.NewRecorder()

	// 根据参数生成请求，各种配置、参数在这里
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)

	// 根据路由 router 构建 handleFunc(w http.ResponseWriter, req *http.Request), 起服务
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}

// TestUserSave router("/user?name=xiaotaot") get 默认值
func TestUserSaveQueryNoAge(t *testing.T) {
	username := "xiaotaot"
	router := initRouter.SetupRouter()

	// 生成一条 httptest 的请求记录，NewRecorder() 是 http.ResponseWriter 的实现
	w := httptest.NewRecorder()

	// 根据参数生成请求，各种配置、参数在这里
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)

	// 根据路由 router 构建 handleFunc(w http.ResponseWriter, req *http.Request), 起服务
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+",年龄:20已经保存", w.Body.String())
}
