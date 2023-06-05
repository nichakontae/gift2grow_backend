package profile

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
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

	if result := mysql.Gorm.Where("user_id = ?", userId).Find(&donate_history); result.Error != nil {
		return &response.GenericError{
			Message: "User not found",
			Err:     nil,
		}
	}

	for i := 0; i < len(donate_history); i++ {
		var details model.Campaign
		if result := mysql.Gorm.Where("Id = ?", donate_history[i].CampaignId).First(&details); result.Error != nil {
			return &response.GenericError{
				Message: "Unable to get detail of this campaign",
				Err:     result.Error,
			}
		}
		donate_history[i].Campaign = &details
	}

	return c.JSON(donate_history)
}
