package main

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"work/internal/config"
	"work/internal/pkg/log"
)

func main() {
	Init()

	srv := &http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: config.IrisInit(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			zap.L().Error("Server ListenAndServe", zap.Error(err))
			panic(err)
		}
	}()

	zap.L().Debug("server run on debug mode")

	closeServer(srv)
}

func Init() {
	config.Init()

	log.Init()

	config.MemorySQLiteInit()
	config.IrisInit()
}

func closeServer(srv *http.Server) {
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			panic(err)
		}
	}(zap.L())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exited")

}
