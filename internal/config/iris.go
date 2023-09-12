package config

import (
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"work/internal/consts"
)

var (
	router iris.Party
)

func newLogFile() *os.File {
	f, err := os.OpenFile(consts.DefaultIrisLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func IrisInit() *iris.Application {
	app := iris.New()
	app.Logger().SetOutput(io.MultiWriter(newLogFile(), os.Stdout))
	app.Configure(iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05",
		Charset:                           "UTF-8",
	}))
	router = app.Party("/api")
	return app
}

func R() iris.Party {
	return router
}
