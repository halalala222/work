package main

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"work/internal/config"
	doctordelivery "work/internal/doctor/delivery"
	doctorrepo "work/internal/doctor/repository"
	doctoruc "work/internal/doctor/usecase"
	drugdelivery "work/internal/drug/delivery"
	drugrepo "work/internal/drug/repository"
	druguc "work/internal/drug/usecase"
	"work/internal/pkg/log"
	saledelivery "work/internal/sales/delivery"
	salesrepo "work/internal/sales/repository"
	salesuc "work/internal/sales/usecase"
)

func main() {
	Init()
	app := config.FiberInit()
	api := app.Group("/api")

	ctx := context.Background()
	db := config.DB(ctx)

	zap.L().Debug("server run on debug mode")

	drugRepo := drugrepo.NewDrugRepo(db)
	drugUseCase := druguc.NewDrugUseCase(drugRepo)
	drugDelivery := drugdelivery.NewDrugDelivery(drugUseCase)
	drugdelivery.DrugRouter(api, drugDelivery)

	doctorRepo := doctorrepo.NewDoctorRepo(db)
	doctorUseCase := doctoruc.NewDoctorUseCase(doctorRepo)
	doctorDelivery := doctordelivery.NewDoctorDelivery(doctorUseCase)
	doctordelivery.DoctorRouter(api, doctorDelivery)

	salesRepo := salesrepo.NewSalesRepo(db)
	salesUseCase := salesuc.NewSalesUseCase(salesRepo)
	salesDelivery := saledelivery.NewSalesDelivery(salesUseCase)
	saledelivery.SalesRouter(api, salesDelivery)

	zap.L().Fatal("listen error", zap.Error(app.Listen(viper.GetString("server.addr"))))
}

func Init() {
	config.Init()

	log.Init()

	config.MemorySQLiteInit()
}
