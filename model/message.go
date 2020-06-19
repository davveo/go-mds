package model

import (
	"sync/atomic"
)

type Message struct {
	Id               int64  `xorm:"pk autoincr"`                            // 主键
	Status           string `xorm:"varchar(20) notnull default('')"`        // 状态
	Creater          string `xorm:"varchar(100) default('')"`               // 创建者
	Editor           string `xorm:"varchar(100) default('')"`               // 编辑者
	MessageId        string `xorm:"varchar(50) unique notnull default('')"` // 消息ID
	MessageBody      string `xorm:"varchar(255) notnull default('')"`       // 消息体
	Remark           string `xorm:"varchar(200) default('')"`               // 备注
	MessageSendTimes int32  `xorm:"int notnull"`                            // 消息重发次数
	MessageDataType  string `xorm:"varchar(50)  null default('')"`          // 消息数据类型
	ConsumerQueue    string `xorm:"varchar(100) notnull default('')"`       // 消费队列
	AreadlyDead      string `xorm:"varchar(20) notnull default('')"`        // 是否死亡
	Version          int    `xorm:"version"`                                //版本
	CreatedTime      int    `xorm:"created"`                                //创建时间
	DeletedTime      int    `xorm:"deleted"`                                //删除时间
	UpdatedTime      int    `xorm:"updated"`                                //修改后自动更新时间
}

func NewMessage(messageId string, messageBody string, consumerQueue string) *Message {
	return &Message{MessageId: messageId, MessageBody: messageBody, ConsumerQueue: consumerQueue}
}

func (this *Message) GetMessageId() string {
	return this.MessageId
}

func (this *Message) SetMessageId(messageId string) {
	this.MessageId = messageId
}

func (this *Message) GetStatus() string {
	return this.Status
}

func (this *Message) SetStatus(status string) {
	this.Status = status
}

func (this *Message) GetMessageBody() string {
	return this.MessageBody
}

func (this *Message) SetMessageBody(messageBody string) {
	this.MessageBody = messageBody
}

func (this *Message) GetMessageDataType() string {
	return this.MessageDataType
}

func (this *Message) SetMessageDataType(messageDataType string) {
	this.MessageDataType = messageDataType
}

func (this *Message) GetConsumerQueue() string {
	return this.ConsumerQueue
}

func (this *Message) SetConsumerQueue(consumerQueue string) {
	this.ConsumerQueue = consumerQueue
}

func (this *Message) GetAreadlyDead() string {
	return this.AreadlyDead
}

func (this *Message) SetAreadlyDead(areadlyDead string) {
	this.AreadlyDead = areadlyDead
}

func (this *Message) GetMessageSendTimes() int32 {
	return this.MessageSendTimes
}

func (this *Message) SetMessageSendTimes(messageSendTimes int32) {
	this.MessageSendTimes = messageSendTimes
}

func (this *Message) AddSendTimes() int32 {
	return atomic.AddInt32(&this.MessageSendTimes, 1)
}
