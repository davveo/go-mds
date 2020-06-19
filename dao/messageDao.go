package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	M "github.com/zbrechave/go-mds/model"
	"github.com/zbrechave/go-mds/utils/public"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123123", "127.0.0.1", "3306", "test"))
	public.CheckError(err)
	engine.ShowSQL(true)
	_ = engine.Sync2(new(M.Message))
}

type MessageDao struct{}

func (md *MessageDao) Insert(message *M.Message) (int64, error) {
	return engine.Insert(message)
}

func (md *MessageDao) Update(message *M.Message) (int64, error) {
	return engine.ID(message.Id).Update(message)
}

func (md *MessageDao) Delete(message *M.Message) (int64, error) {
	return engine.Delete(message)
}

func (md *MessageDao) Get(message *M.Message) (bool, error) {
	return engine.Get(message)
}

func (md *MessageDao) GetMessageByMessageId(messageId string) (bool, *M.Message) {
	message := &M.Message{}
	has, _ := engine.Where("message_id = ?", messageId).Get(message)
	return has, message
}

func (md *MessageDao) DeleteMessageByMessageId(messageId string) (int64, error) {
	return engine.Where("message_id = ?", messageId).Delete(&M.Message{})
}
