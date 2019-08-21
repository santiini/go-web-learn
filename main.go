package main

import (
	"flag"
	"fmt"
	"go-web/demo1"
	"go-web/demo2"
	"go-web/demo3"
	"go-web/demo4"
)

// 程序入口函数 main
func main() {
	var demoType string
	// 声明
	flag.StringVar(&demoType, "demo", "all", "whitch demo")
	// 解析
	flag.Parse()

	fmt.Println(demoType)
	if demoType == "demo1" {
		// 测试 json
		demo1.TestDemo11()
	}

	if demoType == "demo2" {
		// 测试 web api
		demo2.TestDemo21()
	}

	if demoType == "demo3" {
		// web api 测试： form 提交，文件上传
		demo3.TestDemo31()
	}
	if demoType == "demo4" {
		// web api: db 数据库操作
		demo4.TestDemo41()
	}

}
