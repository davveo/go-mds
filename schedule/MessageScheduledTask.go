package schedule

import (
	"log"

	"github.com/sirupsen/logrus"
)

/*
	消息处理定时器
		1.处理状态为“待确认”但已超时的消息
		2.处理状态为“发送中”但超时没有被成功消费确认的消息
*/

func Init() {
	log.Println("定时任务开始执行>>>")

	// 处理状态为“待确认”但已超时的消息
	go handleWaitingConfirmTimeOutMessages()

	// 处理状态为“发送中”但超时没有被成功消费确认的消息
	go handleSendingTimeOutMessage()
}

func handleWaitingConfirmTimeOutMessages() {
	logrus.Info("执行(处理[waiting_confirm]状态的消息)任务开始")
	// TODO
	logrus.Info("执行(处理[waiting_confirm]状态的消息)任务结束")
}

func handleSendingTimeOutMessage() {
	logrus.Info("执行(处理[SENDING]的消息)任务开始")
	// TODO
	logrus.Info("执行(处理[SENDING]的消息)任务结束")
}
