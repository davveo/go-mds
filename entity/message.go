package entity

type MessageEntity struct {
	MessageId        string // 消息ID
	MessageBody      string // 消息体
	MessageSendTimes int    // 消息重发次数
	MessageDataType  string // 消息数据类型
	ConsumerQueue    string // 消费队列
	AreadlyDead      string // 是否死亡
	Status           string // 消息状态
	Field1           string // 扩展字段1
	Field2           string // 扩展字段2
	Field3           string // 扩展字段3
}

func (this *MessageEntity) Message(messageId, messageBody, consumerQueue string) {
	this.MessageId = messageId
	this.MessageBody = messageBody
	this.ConsumerQueue = consumerQueue
}

func (this *MessageEntity) GetMessageId() string {
	return this.MessageId
}

func (this *MessageEntity) SetMessageId(messageId string) {
	this.MessageId = messageId
}

func (this *MessageEntity) GetStatus() string {
	return this.Status
}

func (this *MessageEntity) SetStatus(status string) {
	this.Status = status
}

func (this *MessageEntity) GetMessageBody() string {
	return this.MessageBody
}

func (this *MessageEntity) SetMessageBody(messageBody string) {
	this.MessageBody = messageBody
}

func (this *MessageEntity) GetMessageDataType() string {
	return this.MessageDataType
}

func (this *MessageEntity) SetMessageDataType(messageDataType string) {
	this.MessageDataType = messageDataType
}

func (this *MessageEntity) GetConsumerQueue() string {
	return this.ConsumerQueue
}

func (this *MessageEntity) SetConsumerQueue(consumerQueue string) {
	this.ConsumerQueue = consumerQueue
}

func (this *MessageEntity) GetAreadlyDead() string {
	return this.AreadlyDead
}

func (this *MessageEntity) SetAreadlyDead(areadlyDead string) {
	this.AreadlyDead = areadlyDead
}

func (this *MessageEntity) GetMessageSendTimes() int {
	return this.MessageSendTimes
}

func (this *MessageEntity) SetMessageSendTimes(messageSendTimes int) {
	this.MessageSendTimes = messageSendTimes
}

func (this *MessageEntity) GetField1() string {
	return this.Field1
}

func (this *MessageEntity) SetField1(field1 string) {
	this.Field1 = field1
}

func (this *MessageEntity) GetField2() string {
	return this.Field2
}

func (this *MessageEntity) SetField2(field2 string) {
	this.Field2 = field2
}

func (this *MessageEntity) GetField3() string {
	return this.Field3
}

func (this *MessageEntity) SetField3(field3 string) {
	this.Field3 = field3
}

func (this *MessageEntity) AddSendTimes() string {
	this.MessageSendTimes++
}
