package profile

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func EditProfile(c *fiber.Ctx) error {
	var user model.User

	userId := c.Locals("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "User id not found",
			Err:     nil,
		}
	}

	if result := mysql.Gorm.Where("user_id = ?", userId).First(&user); result.Error != nil {
		return &response.GenericError{
			Message: "User not found",
			Err:     nil,
		}
	}

	if err := c.BodyParser(&user); err != nil {
		return &response.GenericError{
			Message: "Error parsing body",
			Err:     err,
		}
	}

	if result := mysql.Gorm.Save(&user); result.Error != nil {
		return &response.GenericError{
			Message: "Error saving user",
			Err:     result.Error,
		}
	}

	return c.JSON(user)
}
