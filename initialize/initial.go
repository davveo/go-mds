package initialize

import (
	"github.com/zbrechave/go-mds/initialize/conf"
	"github.com/zbrechave/go-mds/initialize/cron"
	"github.com/zbrechave/go-mds/initialize/db"
	"github.com/zbrechave/go-mds/initialize/log"
	"github.com/zbrechave/go-mds/initialize/mq"
	"github.com/zbrechave/go-mds/initialize/redis"
	"github.com/zbrechave/go-mds/utils/public"
)

func Init() {
	var err error
	// 先加载配置
	err = conf.Init()
	public.CheckError(err)

	// 数据库健康检测
	err = cron.Init()
	public.CheckError(err)

	// 日志
	err = log.Init()
	public.CheckError(err)

	// 再初始化各个组件
	err = redis.Init()
	public.CheckError(err)

	// 数据初始化
	err = db.Init()
	public.CheckError(err)

	// mq初始化
	err = mq.Init()
	public.CheckError(err)

}
