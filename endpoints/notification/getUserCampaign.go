package notification

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetUserCampaign(c *fiber.Ctx) error {

	campaignId := c.Query("campaignId")
	if campaignId == "" {
		return &response.GenericError{
			Message: "campaignId is missing from query parameters",
			Err:     nil,
		}
	}

	var userCampaign []model.DonateHistory

	if result := mysql.Gorm.Where("campaign_id = ?", campaignId).Find(&userCampaign); result.Error != nil {
		return &response.GenericError{
			Message: "Error when query user campaign",
			Err:     result.Error,
		}
	}

	if len(userCampaign) == 0 {
		return &response.GenericError{
			Message: "No user campaign found",
			Err:     nil,
		}
	}

	userIds := make([]string, len(userCampaign))
	for i, v := range userCampaign {
		if v.UserId != nil {
			userIds[i] = *v.UserId
		}
	}

	return c.JSON(userIds)
}
