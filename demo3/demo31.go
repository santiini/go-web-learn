package demo3

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

// TestDemo31 表单输入
func TestDemo31() {
	http.HandleFunc("/", sayHi)
	http.HandleFunc("/login", login)
	http.HandleFunc("/form", form)
	http.HandleFunc("/upload", upload)

	fmt.Println("server is start at http://localhost:8009")

	err := http.ListenAndServe(":8009", nil)
	// err := http.ListenAndServe(":8009", &listenHandler{})

	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

type listenHandler struct{}

func (a listenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server start at http://localhost:8009")
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	// 解析参数
	r.ParseForm()

	if r.URL.Path != "/" {
		// http 的错误封装
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	fmt.Println(r.Form)

	json.NewEncoder(w).Encode("Hello Golang")
}

// login 表单提交
func login(w http.ResponseWriter, r *http.Request) {
	// 请求的方法
	fmt.Println("method: ", r.Method)
	fmt.Println("form: ", r.Form)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		// 参数校验
		getint, err := strconv.Atoi(r.Form.Get("age"))

		if err != nil {
			fmt.Println("age is not number")
		}

		if getint > 100 {
			fmt.Println("age is too large")
		}

		// 取值方式1： 调用 r.FormValue() 会自动调用 r.ParseForm
		fmt.Println("username: ", r.FormValue("username"))
		fmt.Println("password: ", r.FormValue("password"))

		// 取值方式2： 请求的是登陆数据，那么执行登陆的逻辑判断
		// fmt.Println("username: ", r.Form["username"])
		// fmt.Println("password: ", r.Form["password"])

		// 取值方式3： r.Form.Get("age")
		fmt.Println("username: ", r.Form.Get("username"))
		fmt.Println("password: ", r.Form.Get("password"))
		fmt.Println("age: ", r.Form.Get("age"))
	}
}

// form 防止多次递交表单
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		// ParseForm 解析
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		curtime := time.Now().Unix()
		h := md5.New()

		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		// 渲染 html
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)

		// 渲染 html
		// http.ServeFile(w, r, "form.html")
	} else {
		// 登录
		// get请求就解析url传递的参数，POST则解析响应包的主体
		// r.ParseForm()

		// ParseForm 解析, x-wwww-form-urlencoded 格式的数据
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Println(r.Form)

		token := r.Form.Get("token")
		fmt.Println("token: ", token)

		if token != "" {
			// 验证 token
			fmt.Println("token is right")
		} else {
			fmt.Println("token is invalid")
		}

		fmt.Println("username length", len(r.FormValue("username")))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

/*
	上传文件经过如下 3 步：
		1. 在 Web 页面选择一个文件，然后上传
		2. 在服务端读取上传文件的数据（字节流
		3. 将文件数据写到服务端的某一个文件中

	通过下面的实例我们可以看到我们上传文件主要三步处理：
		1. 表单中增加enctype="multipart/form-data"
		2. 服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
		3. 使用r.FormFile获取文件句柄，然后对文件进行存储等处理。

*/
func upload(w http.ResponseWriter, r *http.Request) {
	// tips: 获取其他非文件字段信息的时候就不需要调用r.ParseForm，因为在需要的时候Go自动会去调用。
	//		而且ParseMultipartForm调用一次之后，后面再次调用不会再有效果。
	// r.ParseForm();
	if r.Method != "POST" {
		http.Error(w, "request method is not support ", http.StatusBadRequest)
	}

	// r.ParseMultipartForm, 参数表示 maxMemory
	// 服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
	// r.ParseMultipartForm(32 << 20)
	r.ParseMultipartForm(1024 * 1024) // 最多在内存中一次处理 1 MB 的数据
	// 通过 r.FormFile 获取上面的文件句柄，
	file, handler, err := r.FormFile("uploadfile")

	if err != nil {
		fmt.Println(err)
		return
	}

	// 延迟关闭文件（在uploadFile函数结束时关闭文件）
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// 使用了io.Copy来存储文件。
	io.Copy(f, file)
}
