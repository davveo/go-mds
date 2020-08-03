package db

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	C "github.com/zbrechave/go-mds/initialize/conf"
)

var (
	Engine *xorm.Engine
)

func Init() error {
	var err error
	Engine, err = xorm.NewEngine("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8",
		C.GetSettings().Mysql.User,
		C.GetSettings().Mysql.Password,
		C.GetSettings().Mysql.DbName))
	if err != nil {
		return err
	} else {
		Engine.SetMaxIdleConns(C.GetSettings().Mysql.MaxIdleConn)
		Engine.SetMaxOpenConns(C.GetSettings().Mysql.MaxOpenConn)
		Engine.ShowSQL(C.GetSettings().Mysql.ShowSQL)
		logrus.Info("数据库初始化成功")
	}

	timer := time.NewTicker(time.Minute * 15)
	go func(engine *xorm.Engine) {
		for range timer.C {
			err = engine.Ping()
			if err != nil {
				logrus.Errorf("db connect error: %#v\n", err.Error())
			} else {
				logrus.Info("数据健康检测正常")
			}
		}
	}(Engine)
	return nil
}
