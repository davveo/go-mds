package impl

import (
	"github.com/zbrechave/go-mds/service/entity"
	"github.com/zbrechave/go-mds/service/enums"
)

type MessageService struct {
}

/**
 * 预存储消息.
 */
func (messageservice *MessageService) saveMessageWaitingConfirm(message entity.Message) string {

	return ""
}

/**
 * 确认并发送消息.
 */
func (messageservice *MessageService) confirmAndSendMessage(messageId string) error {
	return nil
}

/**
 * 存储并发送消息.
 */

func (messageservice *MessageService) saveAndSendMessage(message entity.Message) error {
	return nil
}

/**
 * 直接发送消息.
 */

func (messageservice *MessageService) directSendMessage(message entity.Message) error {
	return nil
}

/**
 * 重发消息.
 */

func (messageservice *MessageService) reSendMessage(message entity.Message) error {
	return nil
}

/**
 * 根据messageId重发某条消息.
 */

func (messageservice *MessageService) reSendMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 将消息标记为死亡消息.
 */

func (messageservice *MessageService) setMessageToAreadlyDead(messageId string) error {
	return nil
}

/**
 * 根据消息ID获取消息
 */

func (messageservice *MessageService) getMessageByMessageId(messageId string) (entity.Message, error) {
	return entity.Message{}, nil
}

/**
 * 根据消息ID删除消息
 */

func (messageservice *MessageService) deleteMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 重发某个消息队列中的全部已死亡的消息.
 */

func (messageservice *MessageService) reSendAllDeadMessageByQueueName(queueName string, batchSize int) error {
	return nil
}
