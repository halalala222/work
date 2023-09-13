package main

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"work/internal/config"
	"work/internal/drug/delivery"
	"work/internal/drug/repository"
	"work/internal/drug/usecase"
	"work/internal/pkg/log"
)

func main() {
	Init()
	app := config.FiberInit()
	api := app.Group("/api")

	ctx := context.Background()
	db := config.DB(ctx)

	zap.L().Debug("server run on debug mode")

	drugRepo := repository.NewDrugRepo(db)
	drugUseCase := usecase.NewDrugUseCase(drugRepo)
	drugDelivery := delivery.NewDrugDelivery(drugUseCase)
	delivery.DrugRouter(api, drugDelivery)
	zap.L().Fatal("listen error", zap.Error(app.Listen(viper.GetString("server.addr"))))
}

func Init() {
	config.Init()

	log.Init()

	config.MemorySQLiteInit()
}
