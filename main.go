package main

import (
	"log"

	"github.com/zbrechave/go-mds/schedule"

	"github.com/zbrechave/go-mds/router"
)

func main() {
	// 装载路由
	r := router.NewRouter()

	schedule.Init()
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
