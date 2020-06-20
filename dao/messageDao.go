package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zbrechave/go-mds/initialize/db"
	M "github.com/zbrechave/go-mds/model"
)

type MessageDao struct{}

func (md *MessageDao) Insert(message *M.Message) (int64, error) {
	return db.Engine.Insert(message)
}

func (md *MessageDao) Update(message *M.Message) (int64, error) {
	return db.Engine.ID(message.Id).Update(message)
}

func (md *MessageDao) Delete(message *M.Message) (int64, error) {
	return db.Engine.Delete(message)
}

func (md *MessageDao) Get(message *M.Message) (bool, error) {
	return db.Engine.Get(message)
}

func (md *MessageDao) GetMessageList(size, page int) (message []M.Message, err error) {
	err = db.Engine.Limit(size, page).Find(&message)
	return message, err
}

func (md *MessageDao) GetMessageByMessageId(messageId string) (bool, *M.Message) {
	message := &M.Message{}
	has, _ := db.Engine.Where("message_id = ?", messageId).Get(message)
	return has, message
}

func (md *MessageDao) DeleteMessageByMessageId(messageId string) (int64, error) {
	return db.Engine.Where("message_id = ?", messageId).Delete(&M.Message{})
}
