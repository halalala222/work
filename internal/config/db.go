package config

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"
	"work/internal/consts"
	"work/internal/domain"
)

var (
	db *gorm.DB
)

func MemorySQLiteInit() {
	var err error
	if db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger:      getGormLogger(),
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(
		&domain.Drug{},
		&domain.Doctor{},
		&domain.Sales{},
	); err != nil {
		panic(err)
	}
}

func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}

func getGormLogger() logger.Interface {
	ignoreRecordNotFound := false
	logLevel := logger.Info
	if !viper.GetBool("debug") {
		ignoreRecordNotFound = true
		logLevel = logger.Error
	}
	logFile, err := os.OpenFile(consts.DefaultSQLLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	multiOutput := io.MultiWriter(os.Stdout, logFile)
	return logger.New(
		log.New(multiOutput, "[DB] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: ignoreRecordNotFound,
			Colorful:                  false,
		},
	)
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("get sqlDB error", zap.Any("error", err))
	}
	if err = sqlDB.Close(); err != nil {
		zap.L().Error("close sqlDB error", zap.Any("error", err))
	}
}
