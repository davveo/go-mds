package model

import (
	"sync/atomic"
)

type Message struct {
	Id               int64  `xorm:"pk autoincr" json:"id"`                                           // 主键
	Status           string `xorm:"varchar(20) notnull default('')" json:"status"`                   // 状态
	Creator          string `xorm:"varchar(100) default('')" json:"creator,omitempty"`               // 创建者
	MessageId        string `xorm:"varchar(50) unique notnull default('')" json:"messageid"`         // 消息ID
	Editor           string `xorm:"varchar(100) default('')" json:"editor,omitempty"`                // 编辑者
	MessageBody      string `xorm:"varchar(255) notnull default('')" json:"messagebody"`             // 消息体
	Remark           string `xorm:"varchar(200) default('')" json:"editor,omitempty"`                // 备注
	MessageDataType  string `xorm:"varchar(50)  null default('')" json:"messagedatatype,omitempty"`  // 消息数据类型
	ConsumerQueue    string `xorm:"varchar(100) notnull default('')" json:"consumerqueue,omitempty"` // 消费队列
	AreadlyDead      string `xorm:"varchar(20) notnull default('')" json:"areadlydead,omitempty"`    // 是否死亡
	Version          int    `xorm:"version" json:"version,omitempty"`                                //版本
	DeletedTime      int    `xorm:"deleted" json:"deletedtime,omitempty"`                            //删除时间
	MessageSendTimes int32  `xorm:"int notnull" json:"messagesendtimes,omitempty"`                   // 消息重发次数
	CreatedTime      int    `xorm:"created" json:"createdtime,omitempty"`                            //创建时间
	UpdatedTime      int    `xorm:"updated" json:"updatedtime,omitempty"`                            //修改后自动更新时间
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
