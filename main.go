package main

import (
	"fmt"
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/zbrechave/go-mds/initialize"
	C "github.com/zbrechave/go-mds/initialize/conf"
	"github.com/zbrechave/go-mds/router"
	"github.com/zbrechave/go-mds/schedule"
)

func init() {

}

func main() {
	gin.SetMode(gin.DebugMode)

	// 装载路由
	routers := router.NewRouter()

	// 配置初始化
	initialize.Init()

	// 异步任务初始化
	schedule.Init()

	endless.DefaultReadTimeOut = C.GetSettings().Server.ReadTimeOut
	endless.DefaultWriteTimeOut = C.GetSettings().Server.WriteTimeOut
	endless.DefaultMaxHeaderBytes = C.GetSettings().Server.MaxHeaderBytes
	endPoint := fmt.Sprintf(":%d", C.GetSettings().Server.Port)

	server := endless.NewServer(endPoint, routers)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server err: %v", err)
	}
}
