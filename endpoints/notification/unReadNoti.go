package notification

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func UnreadNoti(c *fiber.Ctx) error {
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "User id not found",
			Err:     nil,
		}
	}
	var user_noti []model.UserNoti

	if result := mysql.Gorm.Where("user_id = ? AND is_read = ?", userId, false).Find(&user_noti); result.Error != nil {
		return &response.GenericError{
			Message: "user not found",
			Err:     nil,
		}
	}

	return c.JSON(len(user_noti))
}
