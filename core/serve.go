package core

import (
	"context"
	"jtyl_feishu_mt/controller"
	"jtyl_feishu_mt/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controller.Ping)

	r.POST("/zntz/out", controller.ZntzOut)
	r.POST("/zntz/in", controller.ZntzIn)
	r.POST("/zntz/data/return", controller.ZntzDataReturn)

	return r
}

func RunServe() {

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: NewRouter(),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			service.Logger.Fatalf("HTTP服务异常: %v", err)
		}
	}()

	<-quit
	service.Logger.Info("正在关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		service.Logger.Errorf("强制关闭: %v", err)
	}

	service.Logger.Info("所有资源已释放")
}
