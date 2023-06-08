package auth

import (
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/config"
	"github.com/gofiber/fiber/v2"
)

func GetKey(c *fiber.Ctx) error {
	return c.JSON(response.InfoResponse{
		Success: true,
		Data:    config.C.JwtSecret,
	})
}
