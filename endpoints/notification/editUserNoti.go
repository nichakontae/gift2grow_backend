package notification

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func EditUserNoti(c *fiber.Ctx) error {
	body := new(payloads.EditNoti)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Error when parsing body",
			Err:     err,
		}
	}

	if body.NotiObjectId == nil {
		return &response.GenericError{
			Message: "notiObjectId is missing from body",
			Err:     nil,
		}
	}

	if body.UserId == nil {
		return &response.GenericError{
			Message: "userId is missing from body",
			Err:     nil,
		}
	}

	var userNoti model.UserNoti

	if result := mysql.Gorm.Where("noti_object_id = ? AND user_id = ?", *body.NotiObjectId, *body.UserId).Find(&userNoti); result.Error != nil {
		return &response.GenericError{
			Message: "Error when querying user noti",
			Err:     result.Error,
		}
	}

	if err := mysql.Gorm.Model(&userNoti).Update("is_read", true).Error; err != nil {
		return &response.GenericError{
			Message: "Error when updating user noti",
			Err:     err,
		}
	}

	return c.JSON(userNoti)
}
