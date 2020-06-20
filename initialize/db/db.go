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

type Db struct {
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
		engine.SetMaxIdleConns(C.GetSettings().Mysql.MaxIdleConn)
		engine.SetMaxOpenConns(C.GetSettings().Mysql.MaxOpenConn)
		engine.ShowSQL(C.GetSettings().Mysql.ShowSQL)
		logrus.Info("数据库初始化成功")
	}
	return nil
}

func (d *Db) GetEngine() *xorm.Engine {
	return engine
}
