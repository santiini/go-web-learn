package main

import (
	initRouter "go-web/initRouter"
)

func main() {
	router := initRouter.SetupRouter()

	// 启动程序，最终调用 http.ListenAndServe
	_ = router.Run(":8010")
}
