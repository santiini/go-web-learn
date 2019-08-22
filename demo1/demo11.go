package demo1

import (
	"encoding/json"
	"fmt"
	"os"
)

// TestDemo11 json, map, struct 的转换
func TestDemo11() {
	test11()
	test12()
	test13()
}

// 返回的 struct
type returnData struct {
	Message string
	Code    int
	Data    interface{}
}

// demo1  struct --> json
func test11() {
	res := returnData{
		Message: "success",
		Code:    200,
		Data:    "aaaaaaa",
	}

	jsonData, err := json.Marshal(res)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("json data")
	// fmt.Println(jsonData)

	os.Stdout.Write(jsonData)
	fmt.Println()
}

// demo2 struct --> json
func test12() {
	type Person struct {
		Fn string
		Ln string
	}
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
		P      Person `json:"Person"`
	}

	per := Person{Fn: "John",
		Ln: "Doe",
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		P:      per,
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
}

// demo3 struct --> json 解析大小写的问题, 使用 struct tag
func test13() {
	type ReturnObj struct {
		Status string      `json:"id"`
		Code   int         `json:"code"`
		Data   interface{} `json:"data"`
	}

	data1 := ReturnObj{
		Status: "success",
		Code:   200,
		Data:   []string{"first", "second", "third"},
	}

	res, err := json.Marshal(data1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data1)
	fmt.Println(res)

	data1.Code = 204
	fmt.Println(data1)
	fmt.Println(res)
}
