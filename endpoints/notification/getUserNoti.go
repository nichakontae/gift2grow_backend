package notification

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetUserNoti(c *fiber.Ctx) error {

	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "dormId is missing from query parameters",
			Err:     nil,
		}
	}

	var userNoti []model.UserNoti

	if result := mysql.Gorm.Preload("User").Preload("NotiObject").Preload("NotiObject.Campaign").Where("user_id = ?", userId).Find(&userNoti); result.Error != nil {
		return &response.GenericError{
			Message: "Error when query user noti",
			Err:     result.Error,
		}
	}

	if len(userNoti) == 0 {
		return &response.GenericError{
			Message: "No user noti found",
			Err:     nil,
		}
	}

	return c.JSON(userNoti)
}
