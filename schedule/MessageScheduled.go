package schedule

type MessageScheduled interface {
	// 处理状态为"待确认"但已超时的消息
	handleWaitingConfirmTimeOutMessages()
	// 处理状态为"发送中"但超时没有被成功消费确认的消息
	handleSendingTimeOutMessage()
}
