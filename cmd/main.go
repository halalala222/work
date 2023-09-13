package main

import (
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"work/internal/config"
	"work/internal/pkg/log"
)

func main() {
	Init()

	srv := config.FiberInit()

	go func() {
		if err := srv.ListenAndServe(viper.GetString("server.addr")); err != nil {
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
}

func closeServer(srv *fasthttp.Server) {
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			panic(err)
		}
	}(zap.L())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(); err != nil {
		zap.L().Error("Server Shutdown", zap.Error(err))
	}
	zap.L().Info("Server exited")

}
