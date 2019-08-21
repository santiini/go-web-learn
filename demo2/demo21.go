package demo2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// TestDemo21 web api
func TestDemo21() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/map", writeMap)
	http.HandleFunc("/format", formatResponse)
	http.HandleFunc("/json", writeJSON)
	// 设置监听端口
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		// log.Fatal 打印错误并退出程序
		log.Fatal("ListenAndServe: ", err)
	}
}

// 1. 返回 string
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析参数，默认不解析
	r.ParseForm()

	// 传送到服务端的数据
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", v)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	// 方式1： fmt 写入 这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "Hello writeContent")

	// 方式2: io 写入 使用 io 包而不是 fmt 包来输出字符串，这样源文件编译成可执行文件后，体积要小很多，运行起来也更省资源。
	// io.WriteString(w, "Hello, world! GoLang")
}

// Res 返回结果
type Res map[string]interface{}

// 返回 json
func writeMap(w http.ResponseWriter, r *http.Request) {
	result := Res{
		"name": "xiaot",
		"age":  10,
	}

	jsonRet, err := json.Marshal(result)

	if err != nil {
		fmt.Fprintf(w, "error")
	}

	w.Write(jsonRet)
}

// Result 格式化 Response
type Result struct {
	status string
	code   int
	data   interface{}
}

// 返回标准结果的指针
func newResult(status string, code int, data interface{}) *Result {
	return &Result{
		status,
		code,
		data,
	}
}

// success Result
func successResult(code int, data interface{}) *Result {
	return &Result{
		"success",
		code,
		data,
	}
}

// error Result
func errorResult(code int, data interface{}) *Result {
	return &Result{
		"error",
		code,
		data,
	}
}

// 返回标准的 json 结果
func formatResponse(w http.ResponseWriter, r *http.Request) {
	// res := successResult(200, []string{"first", "second", "third"})
	// res := successResult(200, "Hello format response")

	res := Result{
		"success",
		200,
		[]string{"first", "second"},
	}

	jsonRet, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRet)

	// json.NewEncoder(w).Encode(jsonRet)
}

// 返回  json
func writeJSON(w http.ResponseWriter, r *http.Request) {

	// 1. 利用 jsonApi 返回 json
	// json.NewEncoder(w).Encode(map[string]string{
	// 	"status": "OK",
	// })

	// 2. 使用 struct
	res := &Result{
		"success",
		200,
		[]string{"first", "second"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
