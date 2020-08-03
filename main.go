package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zbrechave/go-mds/initialize"
	C "github.com/zbrechave/go-mds/initialize/conf"
	"github.com/zbrechave/go-mds/router"
	"github.com/zbrechave/go-mds/schedule"
)

func init() {
	// 配置初始化
	initialize.Init()

	// 异步任务初始化
	schedule.Init()
}

func main() {
	routers := router.NewRouter()
	if err := routers.Run(fmt.Sprintf(":%d", C.GetSettings().Server.Port)); err != nil {
		logrus.Errorf("Server err: %v", err.Error())
	}
}
