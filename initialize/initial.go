package initialize

import (
	"github.com/zbrechave/go-mds/initialize/conf"
	"github.com/zbrechave/go-mds/initialize/cron"
	"github.com/zbrechave/go-mds/initialize/db"
	"github.com/zbrechave/go-mds/initialize/log"
	"github.com/zbrechave/go-mds/initialize/migrate"
	"github.com/zbrechave/go-mds/initialize/mq"
	"github.com/zbrechave/go-mds/initialize/redis"
	"github.com/zbrechave/go-mds/utils/public"
)

func Init() {
	var err error
	// 配置初始化
	err = conf.Init()
	public.CheckError(err)

	// 日志初始化
	err = log.Init()
	public.CheckError(err)

	// Redis初始化
	err = redis.Init()
	public.CheckError(err)

	// 数据库初始化
	err = db.Init()
	public.CheckError(err)

	// mq初始化
	err = mq.Init()
	public.CheckError(err)

	// 数据库健康检测
	err = cron.Init()
	public.CheckError(err)

	// 数据库迁移
	err = migrate.Init()
	public.CheckError(err)

}
