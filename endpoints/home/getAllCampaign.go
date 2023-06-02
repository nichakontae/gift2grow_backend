package home

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/config"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetAllCampaign(c *fiber.Ctx) error {
	var campaigns []model.Campaign

	if result := mysql.Gorm.Find(&campaigns); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all campaign",
			Err:     result.Error,
		}
	}
	var donateHistory model.DonateHistory
	CampaignPayloads := make([]*payloads.AllCampaign, 0)

	for i := 0; i < len(campaigns); i++ {
		var count int64
		coverImage, _ := url.JoinPath(config.C.ProductionURL, *campaigns[i].CoverImage)
		campaigns[i].CoverImage = &coverImage

		if result := mysql.Gorm.Model(&donateHistory).Where("id = ?", campaigns[i].Id).Count(&count); result.Error != nil {
			return &response.GenericError{
				Message: "Unable to get danate history",
				Err:     result.Error,
			}
		}

		campaignPayload := &payloads.AllCampaign{
			CampaignId:      campaigns[i].Id,
			CoverImage:      campaigns[i].CoverImage,
			SchoolName:      campaigns[i].SchoolName,
			IsCompleted:     campaigns[i].IsCompleted,
			CompletedAmount: campaigns[i].CompletedAmount,
			TrackingAmount:  &count,
			CreatedAt:       campaigns[i].CreatedAt,
		}
		CampaignPayloads = append(CampaignPayloads, campaignPayload)
	}

	return c.JSON(CampaignPayloads)
}
