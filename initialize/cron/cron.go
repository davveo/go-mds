package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/zbrechave/go-mds/initialize/db"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func Init() error {
	c := newWithSeconds()
	_, err := c.AddFunc("*/5 * * * * ?", db.CheckHealth)
	if err != nil {
		logrus.Error("数据健康检测定时任务失败")
		return err
	}
	c.Start()

	return nil
}
