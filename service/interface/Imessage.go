package _interface

import M "github.com/zbrechave/go-mds/model"

type Imessage interface {
	/**
	 * 预存储消息.
	 */
	SaveMessageWaitingConfirm(message M.Message) string

	/**
	 * 确认并发送消息.
	 */
	confirmAndSendMessage(messageId string) error

	/**
	 * 存储并发送消息.
	 */
	saveAndSendMessage(message M.Message) error

	/**
	 * 直接发送消息.
	 */
	directSendMessage(message M.Message) error

	/**
	 * 重发消息.
	 */
	reSendMessage(message M.Message) error

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
	getMessageByMessageId(messageId string) (M.Message, error)

	/**
	 * 根据消息ID删除消息
	 */
	deleteMessageByMessageId(messageId string) error

	/**
	 * 重发某个消息队列中的全部已死亡的消息.
	 */
	reSendAllDeadMessageByQueueName(queueName string, batchSize int) error
}
