package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	data := make(map[string]interface{})
	data["id"] = 1

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1001,
		"msg":  "success",
		"data": data,
	})
}
