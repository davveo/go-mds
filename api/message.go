package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	M "github.com/zbrechave/go-mds/model"
	"github.com/zbrechave/go-mds/model/request"
	"github.com/zbrechave/go-mds/service"
)

func Create(ctx *gin.Context) {
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
		logrus.Fatal(err)
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
