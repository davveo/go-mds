package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/zbrechave/go-mds/entity"
	"github.com/zbrechave/go-mds/utils/public"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine(
		"mysql",
		"root:123123@locolhost:3306/test?charset=utf8")
	public.CheckError(err)

	engine.ShowSQL(true)
}

type MessageDao struct{}

func (md *MessageDao) Insert(message *entity.MessageEntity) (int64, error) {
	return engine.Insert(message)
}

func (md *MessageDao) Update(message *entity.MessageEntity) (int64, error) {
	return engine.Insert(message)
}

func (md *MessageDao) Delete(message *entity.MessageEntity) (int64, error) {
	return engine.Insert(message)
}

func (md *MessageDao) Find(message *entity.MessageEntity) (int64, error) {
	return engine.Insert(message)
}
