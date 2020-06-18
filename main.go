package main

import (
	"log"

	"github.com/zbrechave/go-mds/server"
)

func main() {
	// 装载路由
	r := server.NewRouter()
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
