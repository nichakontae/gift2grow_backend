package notification

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func PostUserNoti(c *fiber.Ctx) error {
	body := new(payloads.PostNoti)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Error when parsing body",
			Err:     err,
		}
	}

	for _, userID := range body.UserIds {
		if userID == "" {
			return &response.GenericError{
				Message: "userId is missing from body",
				Err:     nil,
			}
		} else {
			var user model.User
			if result := mysql.Gorm.Where("id = ?", userID).Find(&user); result.Error != nil {
				return &response.GenericError{
					Message: "Error when querying user",
					Err:     result.Error,
				}
			}
			if user.Id == nil {
				return &response.GenericError{
					Message: "No user found",
					Err:     nil,
				}
			}
		}
	}

	NotiObject := &model.NotiObject{
		CampaignId: body.CampaignId,
		UserNoti:   []*model.UserNoti{},
	}

	isRead := false
	for _, userID := range body.UserIds {
		currentUserID := userID // Create a separate variable to store the current userID
		userNoti := &model.UserNoti{
			UserId: &currentUserID,
			IsRead: &isRead,
		}
		NotiObject.UserNoti = append(NotiObject.UserNoti, userNoti)
	}

	if result := mysql.Gorm.Create(NotiObject); result.Error != nil {
		return &response.GenericError{
			Message: "Error when creating noti object",
			Err:     result.Error,
		}
	}
	return c.JSON(NotiObject)
}
