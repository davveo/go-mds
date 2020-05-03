package impl

import "github.com/zbrechave/go-mds/service/entity"

type message struct {
}

/**
 * 预存储消息.
 */
func (ms *message) saveMessageWaitingConfirm(message entity.Message) string {
	return ""
}

/**
 * 确认并发送消息.
 */
func (ms *message) confirmAndSendMessage(messageId string) error {
	return nil
}

/**
 * 存储并发送消息.
 */

func (ms *message) saveAndSendMessage(message entity.Message) error {
	return nil
}

/**
 * 直接发送消息.
 */

func (ms *message) directSendMessage(message entity.Message) error {
	return nil
}

/**
 * 重发消息.
 */

func (ms *message) reSendMessage(message entity.Message) error {
	return nil
}

/**
 * 根据messageId重发某条消息.
 */

func (ms *message) reSendMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 将消息标记为死亡消息.
 */

func (ms *message) setMessageToAreadlyDead(messageId string) error {
	return nil
}

/**
 * 根据消息ID获取消息
 */

func (ms *message) getMessageByMessageId(messageId string) (entity.Message, error) {
	return entity.Message{}, nil
}

/**
 * 根据消息ID删除消息
 */

func (ms *message) deleteMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 重发某个消息队列中的全部已死亡的消息.
 */

func (ms *message) reSendAllDeadMessageByQueueName(queueName string, batchSize int) error {
	return nil
}
