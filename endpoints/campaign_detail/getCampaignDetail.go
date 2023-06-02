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

func GetCampaignDetail(c *fiber.Ctx) error {
	var details model.Campaign

	campaignId := c.Query("campaignId")
	if campaignId == "" {
		return &response.GenericError{
			Message: "campaignId is missing from query parameters",
			Err:     nil,
		}
	}

	if result := mysql.Gorm.Preload("CampaignImages").Preload("WantLists").Where("Id  = ?", campaignId).First(&details); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get detail of this campaign",
			Err:     result.Error,
		}
	}
	
	coverImage, _ := url.JoinPath(config.C.ProductionURL, *details.CoverImage)
	var images []string
	for _, image := range details.CampaignImages {
		images = append(images, *image.Image)
	} 
	var lists []string
	for _, list := range details.WantLists {
		lists = append(lists, *list.WantItem)
	}
	detailsPayload := payloads.Campaign{
		CampaignId: details.Id,
		CoverImage: &coverImage,
		Topic: details.Topic,
		Location: details.Location,
		SchoolName: details.SchoolName,
		Description: details.Description,
		IsCompleted: details.IsCompleted,
		TelContact: details.TelContact,
		CompletedAmount: details.CompletedAmount,
		CreatedAt: details.CreatedAt,
		WantLists: lists,
		CampaignImages: images,
	}

	return c.JSON(detailsPayload)
}