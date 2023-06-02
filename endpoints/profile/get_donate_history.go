package profile

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetDonateHistory(c *fiber.Ctx) error {
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "User id not found",
			Err:     nil,
		}
	}
	var donate_history []model.DonateHistory

	if result := mysql.Gorm.Where("user_id = ?", userId).First(&donate_history); result.Error != nil {
		return &response.GenericError{
			Message: "User not found",
			Err:     nil,
		}
	}

	if len(donate_history) == 0 {
		return &response.GenericError{
			Message: "Donate history not found",
			Err:     nil,
		}
	}

	donate_history_payload := make([]*payloads.DonateHistory, len(donate_history))
	for i, v := range donate_history {
		donate_history_payload[i] = &payloads.DonateHistory{
			CampaignId: v.CampaignId,
			Campaign: &payloads.Campaign{
				Topic:      v.Campaign.Topic,
				SchoolName: v.Campaign.SchoolName,
				CoverImage: v.Campaign.CoverImage,
			},
			UserId:         v.UserId,
			TrackingNumber: v.TrackingNumber,
			DonationDate:   v.DonationDate,
		}
	}

	return c.JSON(donate_history)
}
