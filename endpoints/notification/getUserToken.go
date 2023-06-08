package notification

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetUserToken(c *fiber.Ctx) error {
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "userId is missing from query parameters",
			Err:     nil,
		}
	}

	var userTokens []model.UserToken

	if result := mysql.Gorm.Where("user_id = ?", userId).Find(&userTokens); result.Error != nil {
		return &response.GenericError{
			Message: "Error getting user token",
			Err:     result.Error,
		}
	}

	if len(userTokens) == 0 {
		return &response.GenericError{
			Message: "No user token found",
			Err:     nil,
		}
	}

	tokens := make([]string, len(userTokens))
	for i, userToken := range userTokens {
		if userToken.Token != nil {
			tokens[i] = *userToken.Token
		}
	}

	return c.JSON(tokens)
}
