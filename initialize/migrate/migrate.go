package migrate

import (
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	"github.com/zbrechave/go-mds/initialize/db"
	M "github.com/zbrechave/go-mds/model"
)

var d db.Db
var err error
var engine *xorm.Engine

func init() {
	engine = d.GetEngine()
}

func Init() error {
	err = engine.Sync2(new(M.Message))
	if err != nil {
		logrus.Error("数据库迁移失败")
		return err
	} else {
		logrus.Info("数据库迁移成功")
	}
	return nil
}
