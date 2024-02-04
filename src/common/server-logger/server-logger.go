package serverlogger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	loggerFormat = `
requestid: ${locals:requestid}
time: ${time}
method: ${method} ${path}:
status: ${status}​
query: ${queryParams}
body: ${body}
headers: ${reqHeaders}
`
	timeFormat = "01/01/2006 15:04:05"
	timeZone   = "Russia/Moscow"
)

func Init(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     loggerFormat,
		TimeFormat: timeFormat,
		TimeZone:   timeZone,
	}))
}
