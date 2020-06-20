package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/zbrechave/go-mds/initialize/db"
	M "github.com/zbrechave/go-mds/model"
)

var engine *xorm.Engine

func init() {
	engine = db.Engine
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
