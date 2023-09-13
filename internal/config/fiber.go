package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"io"
	"os"
	"time"
	"work/internal/consts"
)

var (
	router fiber.Router
)

func newLogFile() *os.File {
	f, err := os.OpenFile(consts.DefaultFiberLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func FiberInit() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "[Fiber] ${ip}:${port} ${time} ${status} - ${latency} ${method} ${path}\\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       io.MultiWriter(newLogFile(), os.Stdout),
	}))

	v1 := app.Group("/vi")
	api := v1.Group("/api")
	router = api
}

func R() fiber.Router {
	return router
}
