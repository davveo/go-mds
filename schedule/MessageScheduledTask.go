package schedule

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/zbrechave/go-mds/service"
)

/*
	消息处理定时器
		1.处理状态为“待确认”但已超时的消息
		2.处理状态为“发送中”但超时没有被成功消费确认的消息
*/

var (
	messageService service.MessageService
)

func Init() {
	go startCron()
}

func startCron() {
	c := cron.New()
	_ = c.AddFunc("*/2 * * * * ?", func() {
		// 处理状态为“待确认”但已超时的消息
		handleWaitingConfirmTimeOutMessages()
	})
	_ = c.AddFunc("*/3 * * * * ?", func() {
		// 处理状态为“发送中”但超时没有被成功消费确认的消息
		handleSendingTimeOutMessage()
	})
	c.Start()
}

func handleWaitingConfirmTimeOutMessages() {
	logrus.Info("执行(处理[waiting_confirm]状态的消息)任务开始")

	logrus.Info("执行(处理[waiting_confirm]状态的消息)任务结束")
}

func handleSendingTimeOutMessage() {
	logrus.Info("执行(处理[SENDING]的消息)任务开始")
	// TODO
	logrus.Info("执行(处理[SENDING]的消息)任务结束")
}
