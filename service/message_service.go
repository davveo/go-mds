package service

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/zbrechave/go-mds/dao"
	M "github.com/zbrechave/go-mds/model"
	"github.com/zbrechave/go-mds/utils/enums"
)

const (
	MESSAGE_MAX_SEND_TIMES = 5
)

type MessageService struct {
	dao *dao.MessageDao
	rbq *RabbitMQ
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

	message.SetStatus(enums.MESSAGE_WAITING_CONFIRM)
	message.SetAreadlyDead(enums.NO)
	message.SetMessageSendTimes(0)

	// 保证发送的幂等
	has, newmessage := messageService.dao.GetMessageByMessageId(message.MessageId)
	if has && newmessage != nil {
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
	has, message := messageService.dao.GetMessageByMessageId(messageId)
	if !has || message == nil {
		return errors.New("根据消息id查找的消息为空")
	}

	// 防止重复发送
	if message.GetStatus() == enums.MESSAGE_SENDING {
		return nil
	}

	message.SetStatus(enums.MESSAGE_SENDING)
	if _, err := messageService.dao.Update(message); err != nil {
		return errors.New("更新消息失败")
	}

	return messageService.rbq.Send(message.GetConsumerQueue(), message.GetMessageBody())
}

/**
 * 存储并发送消息.
 */
func (messageService *MessageService) SaveAndSendMessage(message *M.Message) (int64, error) {
	if message.GetMessageBody() == "" {
		return -1, errors.New("保存的消息为空")
	}

	if message.GetConsumerQueue() == "" {
		return -1, errors.New("消息队列不能为空")
	}

	message.SetStatus(enums.MESSAGE_SENDING)
	message.SetAreadlyDead(enums.NO)
	message.SetMessageSendTimes(0)
	result, _ := messageService.dao.Insert(message)
	if err := messageService.rbq.Send(message.GetMessageBody(), message.GetConsumerQueue()); err != nil {
		logrus.Fatal(err)
	}
	return result, nil

}

/**
 * 直接发送消息.
 */
func (messageService *MessageService) DirectSendMessage(message *M.Message) error {
	if message.GetMessageBody() == "" {
		return errors.New("保存的消息为空")
	}

	if message.GetConsumerQueue() == "" {
		return errors.New("消息队列不能为空")
	}

	return messageService.rbq.Send(message.GetMessageBody(), message.GetConsumerQueue())
}

/**
 * 重发消息.
 */

func (messageService *MessageService) ReSendMessage(message *M.Message) error {
	if message.GetMessageBody() == "" {
		return errors.New("保存的消息为空")
	}

	if message.GetConsumerQueue() == "" {
		return errors.New("消息队列不能为空")
	}

	message.AddSendTimes()
	_, err := messageService.dao.Update(message)
	if err != nil {
		logrus.Fatal(err)
	}

	return messageService.rbq.Send(message.GetMessageBody(), message.GetConsumerQueue())
}

/**
 * 根据messageId重发某条消息. web
 */

func (messageService *MessageService) ReSendMessageByMessageId(messageId string) error {
	has, message := messageService.GetMessageByMessageId(messageId)
	if !has || message == nil {
		return errors.New("根据消息id查找的消息为空")
	}

	if message.GetMessageSendTimes() >= MESSAGE_MAX_SEND_TIMES {
		message.SetAreadlyDead(enums.YES)
	}
	message.SetMessageSendTimes(message.GetMessageSendTimes() + 1)
	_, err := messageService.dao.Update(message)
	if err != nil {
		logrus.Fatal(err)
	}

	return messageService.rbq.Send(message.GetMessageBody(), message.GetConsumerQueue())
}

/**
 * 将消息标记为死亡消息.
 */

func (messageService *MessageService) SetMessageToAreadlyDead(messageId string) error {
	has, message := messageService.dao.GetMessageByMessageId(messageId)
	if !has || message == nil {
		return errors.New("根据消息id查找的消息为空")
	}

	message.SetAreadlyDead(enums.YES)
	if _, err := messageService.dao.Update(message); err != nil {
		logrus.Fatal(err)
	}
	return nil
}

/**
 * 根据消息ID获取消息
 */

func (messageService *MessageService) GetMessageByMessageId(messageId string) (bool, *M.Message) {
	return messageService.dao.GetMessageByMessageId(messageId)
}

/**
 * 根据消息ID删除消息
 */

func (messageService *MessageService) DeleteMessageByMessageId(messageId string) (int64, error) {
	return messageService.dao.DeleteMessageByMessageId(messageId)
}

/**
 * 重发某个消息队列中的全部已死亡的消息.
 */

func (messageService *MessageService) ReSendAllDeadMessageByQueueName(queueName string, batchSize int) error {
	logrus.Info("==>reSendAllDeadMessageByQueueName")
	//pageNum, numPerPage := 1, 1000
	//if batchSize > 0 && batchSize < 100 {
	//	numPerPage = 100
	//} else if batchSize > 100 && batchSize < 5000 {
	//	numPerPage = batchSize
	//} else if batchSize > 5000 {
	//	numPerPage = 5000
	//} else {
	//	numPerPage = 1000
	//}
	//
	//paramMap := make(map[string]interface{})
	//
	//paramMap["consumerQueue"] = queueName
	//paramMap["areadlyDead"] = enums.YES
	//paramMap["listPageSortType"] = "asc"

	//recordList = []{}
	//
	//Map<String, RpTransactionMessage> messageMap = new HashMap<String, RpTransactionMessage>();
	//List<Object> recordList = new ArrayList<Object>();
	//int pageCount = 1;
	//
	//PageBean pageBean = rpTransactionMessageDao.listPage(new PageParam(pageNum, numPerPage), paramMap);
	//recordList = pageBean.getRecordList();
	//if (recordList == null || recordList.isEmpty()) {
	//	log.info("==>recordList is empty");
	//	return;
	//}
	//pageCount = pageBean.getTotalPage();
	//for (final Object obj : recordList) {
	//	final RpTransactionMessage message = (RpTransactionMessage) obj;
	//	messageMap.put(message.getMessageId(), message);
	//}
	//
	//for (pageNum = 2; pageNum <= pageCount; pageNum++) {
	//	pageBean = rpTransactionMessageDao.listPage(new PageParam(pageNum, numPerPage), paramMap);
	//	recordList = pageBean.getRecordList();
	//
	//	if (recordList == null || recordList.isEmpty()) {
	//		break;
	//	}
	//
	//	for (final Object obj : recordList) {
	//		final RpTransactionMessage message = (RpTransactionMessage) obj;
	//		messageMap.put(message.getMessageId(), message);
	//	}
	//}
	//
	//recordList = null;
	//pageBean = null;
	//
	//for (Map.Entry<String, RpTransactionMessage> entry : messageMap.entrySet()) {
	//	final RpTransactionMessage message = entry.getValue();
	//
	//	message.setEditTime(new Date());
	//	message.setMessageSendTimes(message.getMessageSendTimes() + 1);
	//	rpTransactionMessageDao.update(message);
	//
	//	notifyJmsTemplate.setDefaultDestinationName(message.getConsumerQueue());
	//	notifyJmsTemplate.send(new MessageCreator() {
	//		public Message createMessage(Session session) throws JMSException {
	//			return session.createTextMessage(message.getMessageBody());
	//		}
	//	});
	//}
	return nil
}

func (messageService *MessageService) ListPage(size, page int) ([]M.Message, error) {
	return messageService.dao.GetMessageList(size, page)
}
