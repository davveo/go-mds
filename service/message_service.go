package service

import (
	"errors"

	"github.com/zbrechave/go-mds/dao"

	M "github.com/zbrechave/go-mds/model"
)

const (
	MESSAGE_WAITING_CONFIRM = "待确认"
	MESSAGE_SENDING         = "发送中"
	PUBLIC_YES              = "是"
	PUBLIC_NO               = "否"
)

type MessageService struct {
	dao *dao.MessageDao
}

/**
 * 预存储消息.
 */
func (messageService *MessageService) SaveMessageWaitingConfirm(message *M.Message) (int64, error) {
	if message.GetMessageBody() == "" {
		return 0, errors.New("保存消息为空")
	}
	if message.GetConsumerQueue() == "" {
		return 0, errors.New("消息队列不能为空")
	}
	message.SetStatus(MESSAGE_WAITING_CONFIRM)
	message.SetAreadlyDead(PUBLIC_NO)
	message.SetMessageSendTimes(0)

	// 保证发送的幂等
	has, newmessage := messageService.dao.GetByMessageId(message)
	if has {
		newmessage.AddSendTimes()
		return messageService.dao.Update(newmessage)
	} else {
		return messageService.dao.Insert(message)
	}
}

/**
 * 确认并发送消息.
 */
func (messageService *MessageService) ConfirmAndSendMessage(messageId string) error {
	return nil
}

/**
 * 存储并发送消息.
 */

func (messageService *MessageService) SaveAndSendMessage(message M.Message) error {
	return nil
}

/**
 * 直接发送消息.
 */

func (messageService *MessageService) DirectSendMessage(message M.Message) error {
	return nil
}

/**
 * 重发消息.
 */

func (messageService *MessageService) ReSendMessage(message M.Message) error {
	return nil
}

/**
 * 根据messageId重发某条消息.
 */

func (messageService *MessageService) ReSendMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 将消息标记为死亡消息.
 */

func (messageService *MessageService) SetMessageToAreadlyDead(messageId string) error {
	return nil
}

/**
 * 根据消息ID获取消息
 */

func (messageService *MessageService) GetMessageByMessageId(messageId string) (M.Message, error) {
	return M.Message{}, nil
}

/**
 * 根据消息ID删除消息
 */

func (messageService *MessageService) DeleteMessageByMessageId(messageId string) error {
	return nil
}

/**
 * 重发某个消息队列中的全部已死亡的消息.
 */

func (messageService *MessageService) ReSendAllDeadMessageByQueueName(queueName string, batchSize int) error {
	return nil
}
