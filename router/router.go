package router

import (
	"github.com/zbrechave/go-mds/api"
	"github.com/zbrechave/go-mds/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("/api/v1")
	{
		v1.POST("message/create", api.CreateMessage)
		v1.POST("message/confirm", api.ConfirmMessage)
	}

	admin := r.Group("/api/admin")
	{
		admin.DELETE("message", api.DeleteMessage)        // 删除某条消息
		admin.POST("message/list", api.ListMessage)       // 获取消息列表
		admin.POST("message/send", api.SendMessage)       // 发送单条消息
		admin.POST("message/sendAll", api.SendAllMessage) // 发送所有消息
	}

	return r
}
