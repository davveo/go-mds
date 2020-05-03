package main

import "github.com/micro/go-micro/server"

func main() {
	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
