package notification

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func PostUserToken(c *fiber.Ctx) error {

	body := new(payloads.PostUserToken)

	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Error parsing body",
			Err:     err,
		}
	}

	UserToken := &model.UserToken{
		UserId: body.UserId,
		Token:  body.Token,
	}

	if result := mysql.Gorm.Create(UserToken); result.Error != nil {
		return &response.GenericError{
			Message: "Error creating user token",
			Err:     result.Error,
		}
	}

	return c.JSON(UserToken)
}
