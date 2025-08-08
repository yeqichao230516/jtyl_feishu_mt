package core

import (
	"jtyl_feishu_mt/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controller.Ping)
	r.POST("/zntz/out", controller.ZntzOut)
	r.POST("/zntz/in", controller.ZntzIn)

	return r
}
