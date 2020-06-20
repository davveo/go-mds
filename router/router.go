package router

import (
	"github.com/zbrechave/go-mds/api"
	"github.com/zbrechave/go-mds/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.POST("message/list", api.ConfirmMessage)    // ListMessage
		v1.POST("message/send", api.ConfirmMessage)    // ReSendMessageByMessageId
		v1.POST("message/sendAll", api.ConfirmMessage) // ReSendAllDeadMessageByQueueName
	}

	return r
}
