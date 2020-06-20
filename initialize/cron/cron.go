package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	"github.com/zbrechave/go-mds/initialize/db"
)

var d db.Db
var engine *xorm.Engine

func init() {
	engine = d.GetEngine()
}

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func CheckMysqlHealth() {
	if engine != nil {
		err := engine.Ping()
		if err != nil {
			logrus.Error("数据库不健康")
		} else {
			logrus.Info("数据库正常")
		}
	}
}

func Init() error {
	c := newWithSeconds()
	_, err := c.AddFunc("*/5 * * * * ?", CheckMysqlHealth)
	if err != nil {
		logrus.Error("数据健康检测定时任务失败")
		return err
	} else {
		c.Start()
		logrus.Info("定时任务初始化成功")
	}

	return nil
}
