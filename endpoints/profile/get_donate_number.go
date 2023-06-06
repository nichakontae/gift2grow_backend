package profile

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/response"

	"github.com/gofiber/fiber/v2"
)

func GetDonateNumber(c *fiber.Ctx) error {
	Id := c.Query("Id")
	if Id == "" {
		return &response.GenericError{
			Message: "Campaign id not found",
			Err:     nil,
		}
	}
	var donate_history []model.DonateHistory

	if result := mysql.Gorm.Where("campaign_id = ?", Id).Find(&donate_history); result.Error != nil {
		return &response.GenericError{
			Message: "campaign not found",
			Err:     nil,
		}
	}

	return c.JSON(len(donate_history))
}
