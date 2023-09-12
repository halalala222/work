package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"work/internal/consts"
)

func Init() {
	zap.ReplaceGlobals(getZapLogger(consts.DefaultLogFilePath, zap.AddCaller()))
}

func getZapLogger(logfile string, options ...zap.Option) *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			getEncoder(),
			getLogWriteSyncer(logfile),
			getLogLevel(),
		),
		options...,
	)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogLevel() zapcore.Level {
	if viper.GetBool("DEBUG") {
		return zapcore.DebugLevel
	} else {
		return zapcore.ErrorLevel
	}
}

func getLogWriteSyncer(logFilePath string) zapcore.WriteSyncer {
	if len(logFilePath) == 0 {
		return os.Stdout
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	if viper.GetBool("DEBUG") {
		return zapcore.AddSync(io.MultiWriter(os.Stdout, lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
