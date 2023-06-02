package campaign_detail

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/config"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetCompletedCam(c *fiber.Ctx) error {
	var campaign model.Campaign

	campaignId := c.Query("campaignId")

	if result := mysql.Gorm.Preload("EvidenceCampaignImage").Where("campaign_id = ?", campaignId).Find(&campaign); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all campaign",
			Err:     result.Error,
		}
	}
	var EvidenceImgs []*string
	for i := 0; i < len(campaign.EvidenceCampaignImage); i++ {
		EvidenceCampaignImage, _ := url.JoinPath(config.C.ProductionURL, *campaign.EvidenceCampaignImage[i].Image)

		EvidenceImgs[i] = &EvidenceCampaignImage
	}

	completedCampaign := &payloads.CompletedCampaign{
		CampaignId:            campaign.Id,
		CoverImage:            campaign.CoverImage,
		Topic:                 campaign.Topic,
		SchoolName:            campaign.SchoolName,
		Description:           campaign.Description,
		IsCompleted:           campaign.IsCompleted,
		CompletedAmount:       campaign.CompletedAmount,
		EvidenceCampaignImage: EvidenceImgs,
	}

	return c.JSON(completedCampaign)
}
