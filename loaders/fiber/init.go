package fiber

import (
	"gift2grow_backend/endpoints"
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/config"
	"gift2grow_backend/utils/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var app *fiber.App

func Init() {
	app = fiber.New(fiber.Config{
		ServerHeader:  config.C.ServerHeader,
		ErrorHandler:  errorHandler,
		Prefork:       false,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.InfoResponse{
			Success: true,
			Message: "Gift2Grow_API_ROOT",
		})
	})

	apiGroup := app.Group("api")

	endpoints.Init(apiGroup)

	app.Use(notfoundHandler)

	err := app.Listen(config.C.BackendAddress)
	if err != nil {
		logger.Log(logrus.Fatal, err.Error())
	}
}
