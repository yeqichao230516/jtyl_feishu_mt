package core

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServe() {

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: NewRouter(),
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			Logger.Fatalf("HTTP服务异常: %v", err)
		}
	}()

	<-quit
	Logger.Info("正在关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Logger.Errorf("强制关闭: %v", err)
	}

	Logger.Info("所有资源已释放")
}
