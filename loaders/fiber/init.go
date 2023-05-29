package fiber

import "github.com/gofiber/fiber/v2"

var app *fiber.App

func Init() {
	app = fiber.New(fiber.Config{})
}
