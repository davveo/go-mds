package models

import (
	"sync/atomic"
	"time"
)

type Message struct {
	id               int32     //主键id
	version          int       // 版本号默认为0
	status           string    // 状态 PublicStatusEnum
	creater          string    // 创建人
	editor           string    // 修改人
	remark           string    // 描述
	createTime       time.Time // 创建时间
	editTime         time.Time // 修改时间
	messageSendTimes int32
	serialVersionUID int64
	messageId        string
	messageBody      string
	messageDataType  string
	consumerQueue    string
	areadlyDead      string
	field1           string
	field2           string
	field3           string
}

func (m *Message) MessageId() string {
	return m.messageId
}

func (m *Message) SetMessageId(messageId string) {
	m.messageId = messageId
}

func (m *Message) MessageBody() string {
	return m.messageBody
}

func (m *Message) SetMessageBody(messageBody string) {
	m.messageBody = messageBody
}

func (m *Message) MessageDataType() string {
	return m.messageDataType
}

func (m *Message) SetMessageDataType(messageDataType string) {
	m.messageDataType = messageDataType
}

func (m *Message) ConsumerQueue() string {
	return m.consumerQueue
}

func (m *Message) SetConsumerQueue(consumerQueue string) {
	m.consumerQueue = consumerQueue
}

func (m *Message) MessageSendTimes() int32 {
	return m.messageSendTimes
}

func (m *Message) SetMessageSendTimes(messageSendTimes int32) {
	m.messageSendTimes = messageSendTimes
}

func (m *Message) AddMessageSendTimes() {
	atomic.AddInt32(&(m.messageSendTimes), 1)
}

func (m *Message) AreadlyDead() string {
	return m.areadlyDead
}

func (m *Message) SetAreadlyDead(areadlyDead string) {
	m.areadlyDead = areadlyDead
}

func (m *Message) Field1() string {
	return m.field1
}

func (m *Message) SetField1(field1 string) {
	m.field1 = field1
}

func (m *Message) Field2() string {
	return m.field2
}

func (m *Message) SetField2(field2 string) {
	m.field2 = field2
}

func (m *Message) Field3() string {
	return m.field3
}

func (m *Message) SetField3(field3 string) {
	m.field3 = field3
}

func (m *Message) Id() int32 {
	return m.id
}

func (m *Message) SetId(id int32) {
	m.id = id
}

func (m *Message) EditTime() time.Time {
	return m.editTime
}

func (m *Message) SetEditTime(editTime time.Time) {
	m.editTime = editTime
}

func (m *Message) CreateTime() time.Time {
	return m.createTime
}

func (m *Message) SetCreateTime(createTime time.Time) {
	m.createTime = createTime
}

func (m *Message) Remark() string {
	return m.remark
}

func (m *Message) SetRemark(remark string) {
	m.remark = remark
}

func (m *Message) Editor() string {
	return m.editor
}

func (m *Message) SetEditor(editor string) {
	m.editor = editor
}

func (m *Message) Creater() string {
	return m.creater
}

func (m *Message) SetCreater(creater string) {
	m.creater = creater
}

func (m *Message) Status() string {
	return m.status
}

func (m *Message) SetStatus(status string) {
	m.status = status
}

func (m *Message) Version() int {
	return m.version
}

func (m *Message) SetVersion(version int) {
	m.version = version
}

func NewMessage(messageId string, messageBody string, consumerQueue string) *Message {
	return &Message{
		messageId:     messageId,
		messageBody:   messageBody,
		consumerQueue: consumerQueue,
	}
}
