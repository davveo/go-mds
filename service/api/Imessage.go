package api

type RpTransactionMessage struct {
}

type Imessage interface {
	/**
	 * 预存储消息.
	 */
	SaveMessageWaitingConfirm(Message message) string

	/**
	 * 确认并发送消息.
	 */
	confirmAndSendMessage(messageId string) error

	/**
	 * 存储并发送消息.
	 */
	saveAndSendMessage(Message message) error

	/**
	 * 直接发送消息.
	 */
	directSendMessage(Message message) error

	/**
	 * 重发消息.
	 */
	reSendMessage(Message message) error

	/**
	 * 根据messageId重发某条消息.
	 */
	reSendMessageByMessageId(messageId string) error

	/**
	 * 将消息标记为死亡消息.
	 */
	setMessageToAreadlyDead(messageId string) error

	/**
	 * 根据消息ID获取消息
	 */
	getMessageByMessageId(messageId string) (Message, error)

	/**
	 * 根据消息ID删除消息
	 */
	deleteMessageByMessageId(messageId string) error

	/**
	 * 重发某个消息队列中的全部已死亡的消息.
	 */
	reSendAllDeadMessageByQueueName(queueName string, batchSize int) error
}
