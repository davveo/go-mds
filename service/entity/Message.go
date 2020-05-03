package entity

import "sync/atomic"

type Message struct {
	base             *BaseEntity
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

func NewMessage(messageId string, messageBody string, consumerQueue string) *Message {
	return &Message{
		messageId:     messageId,
		messageBody:   messageBody,
		consumerQueue: consumerQueue,
	}
}
