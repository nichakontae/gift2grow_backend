package campaign_detail

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func PutTamboon(c *fiber.Ctx) error{
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}

	var user model.User

	// Find the user in the database
	if result := mysql.Gorm.Where("id = ?", userId).First(&user); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find user in database",
			Err:     result.Error,
		}
	}

	*user.TamboonPoint = *user.TamboonPoint + 50;

	if result := mysql.Gorm.Save(&user); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to update user",
			Err:     result.Error,
		}
	}

	return c.JSON(user)

}