package campaign_detail

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PostTracking(c *fiber.Ctx) error {
	
	body := new(payloads.Tracking)
    if err := c.BodyParser(body); err != nil {
        return &response.GenericError{
            Message: "Unable to parse body",
            Err:     err,
        }
    }

	user := &model.User{}
	if result := mysql.Gorm.First(user, body.UserId); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find user",
			Err:     result.Error,
		}
	}

	campaign := &model.Campaign{}
	if result := mysql.Gorm.First(campaign, body.CampaignId); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to find campaign",
			Err:     result.Error,
		}
	}

	now := time.Now()

    // Format the date and time as a string in the required format for MySQL
    updatedAt := now.Format("2006-01-02 15:04:05")

    // Parse the createAt string into a time.Time value
    updatedAtTime, err := time.Parse("2006-01-02 15:04:05", updatedAt)
    if err != nil {
        log.Println("handle error")
    }

	DonateHistory := &model.DonateHistory{
		CampaignId: body.CampaignId,
		UserId: body.UserId,
		TrackingNumber: body.TrackingNumber,
		DonationDate: &updatedAtTime,
	}

	if result := mysql.Gorm.Create(&DonateHistory); result.Error != nil {
		return &response.GenericError{
            Message: "Unable to add tracking number",
            Err:     result.Error,
        }
	}
	return c.JSON(DonateHistory)

}