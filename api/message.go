package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	M "github.com/zbrechave/go-mds/model"
	"github.com/zbrechave/go-mds/model/request"
	"github.com/zbrechave/go-mds/service"
)

func CreateMessage(ctx *gin.Context) {
	var messageRequest request.MessageRequest
	var messageService service.MessageService

	if err := ctx.ShouldBindJSON(&messageRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1002, "msg": err.Error(), "data": nil})
		return
	}

	message := M.NewMessage(
		messageRequest.MessageId,
		messageRequest.MesageBody,
		messageRequest.MessageQueue)

	if _, err := messageService.SaveMessageWaitingConfirm(message); err != nil {
		logrus.Println(err)
	}

	data := make(map[string]interface{})
	data["id"] = 1
	data["confirm"] = 0

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1001,
		"msg":  "success",
		"data": data,
	})
}

func ConfirmMessage(ctx *gin.Context) {
	var confirmMessageRequest request.MessageIDRequest
	var messageService service.MessageService

	if err := ctx.ShouldBindJSON(&confirmMessageRequest); err != nil {
		logrus.Println(err)
	}

	if err := messageService.ConfirmAndSendMessage(confirmMessageRequest.MessageId); err != nil {
		logrus.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1001,
		"msg":  "success",
		"data": nil,
	})
}

func SendMessage(ctx *gin.Context) {
	// 发送单条消息
}

func SendAllMessage(ctx *gin.Context) {
	// 发送所有消息
}

func DeleteMessage(ctx *gin.Context) {
	// 删除消息
	var (
		messageR       request.MessageIDRequest
		messageService service.MessageService
	)

	if _, err := messageService.DeleteMessageByMessageId(messageR.MessageId); err != nil {
		logrus.Println("删除消息失败")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1001,
		"msg":  "success",
		"data": nil,
	})
}

func ListMessage(ctx *gin.Context) {

}
