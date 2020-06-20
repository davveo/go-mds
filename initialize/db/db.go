package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	C "github.com/zbrechave/go-mds/initialize/conf"
)

var (
	engine *xorm.Engine
)

type D struct {
}

func Init() error {
	var err error
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8",
		C.GetSettings().Mysql.User,
		C.GetSettings().Mysql.Password,
		C.GetSettings().Mysql.DbName))
	if err != nil {
		return err
	} else {
		engine.SetMaxIdleConns(10)
		engine.SetMaxOpenConns(200)
		engine.ShowSQL(true)
	}
	return nil
}

func CheckHealth() {
	if engine != nil {
		err := engine.Ping()
		if err != nil {
			logrus.Error("数据库不健康")
		} else {
			logrus.Info("数据库正常")
		}
	}
}

func (d *D) GetEngine() *xorm.Engine {
	return engine
}
