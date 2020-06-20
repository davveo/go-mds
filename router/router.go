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
		v1.GET("ping", api.Ping)
		v1.GET("message/list", api.ListMessage)
		v1.POST("message/create", api.CreateMessage)
		v1.POST("message/confirm", api.ConfirmMessage)
	}

	return r
}
