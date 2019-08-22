package test

import (
	initRouter "go-web/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexGetRouter(t *testing.T) {
	router := initRouter.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin", w.Body.String())
}

/*
	运行测试用例
	  1. 在项目根目录，执行 go test -v ./test
*/
