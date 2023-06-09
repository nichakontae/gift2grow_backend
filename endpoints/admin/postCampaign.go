package admin

import (
	"fmt"
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/text"
	"image"
	"image/jpeg"
	"os"

	"github.com/gofiber/fiber/v2"
)

func PostCampaign(c *fiber.Ctx) error {

	// Parse the request body into a payloads.Campaign object
	body := new(payloads.Campaign)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	//Cover Image
	fileHeader, err := c.FormFile("CoverImage")
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse cover image file",
			Err:     err,
		}
	}
	file, err := fileHeader.Open()
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse cover image file",
			Err:     err,
		}
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return &response.GenericError{
			Message: "Unable to decode cover image",
			Err:     err,
		}
	}

	fileSalt := *text.GenerateString(text.GenerateStringSet.Num, 6)

	savingFile, err := os.Create("./images/" + fileSalt + ".jpeg")
	if err != nil {
		return &response.GenericError{
			Message: "Unable to create an cover image file",
			Err:     err,
		}
	}
	defer savingFile.Close()

	if err := jpeg.Encode(savingFile, img, nil); err != nil {
		return &response.GenericError{
			Message: "Unable to save an cover image file",
			Err:     err,
		}
	}

	fileName := fmt.Sprintf("/images/%s.jpeg", fileSalt)

	// Create a new Campaign object with the parsed data from the request body
	campaign := model.Campaign{
		CoverImage:      &fileName,
		Topic:           body.Topic,
		SchoolName:      body.SchoolName,
		Location:        body.Location,
		Description:     body.Description,
		IsCompleted:     body.IsCompleted,
		TelContact:      body.TelContact,
		CompletedAmount: body.CompletedAmount,
	}

	// Save the campaign to the database
	if result := mysql.Gorm.Create(&campaign); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create campaign",
			Err:     result.Error,
		}
	}

	// Parse the multipart form data
	list, err := c.MultipartForm()
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse multipart form",
			Err:     err,
		}
	}

	// Retrieve the wantLists field from the multipart form data
	wantedList := list.Value["wantLists"]
	var wantLists []*model.WantList

	// Create and save each want list item to the database
	for _, item := range wantedList {
		wantList := &model.WantList{
			CampaignId: campaign.Id,
			WantItem:   &item, // Assign the address of 'item' to 'WantItem'
		}
		if result := mysql.Gorm.Create(wantList); result.Error != nil {
			return &response.GenericError{
				Message: "Unable to create want list",
				Err:     result.Error,
			}
		}
		wantLists = append(wantLists, wantList)
	}

	// Update the campaign object with the wantLists
	campaign.WantLists = wantLists

	//campaignImages
	form, err := c.MultipartForm()
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse multipart form",
			Err:     err,
		}
	}

	files := form.File["image"]
	var campaignImages []*model.CampaignImage

	for _, file := range files {
		fileHeader, err := file.Open()
		if err != nil {
			return &response.GenericError{
				Message: "Cannot open image",
				Err:     err,
			}
		}
		defer fileHeader.Close()

		img, _, err := image.Decode(fileHeader)
		if err != nil {
			return &response.GenericError{
				Message: "Unable to decode file as image",
				Err:     err,
			}
		}

		fileSalt := *text.GenerateString(text.GenerateStringSet.Num, 6)

		filePath := fmt.Sprintf("./images/%s.jpeg", fileSalt)
		savingFile, err := os.Create(filePath)
		if err != nil {
			return &response.GenericError{
				Message: "Unable to create an image file",
				Err:     err,
			}
		}
		defer savingFile.Close()

		if err := jpeg.Encode(savingFile, img, nil); err != nil {
			return &response.GenericError{
				Message: "Unable to save an image file",
				Err:     err,
			}
		}

		fileName := fmt.Sprintf("/images/%s.jpeg", fileSalt)
		campaignImage := &model.CampaignImage{
			CampaignId: campaign.Id,
			Image:      &fileName,
		}
		if result := mysql.Gorm.Create(campaignImage); result.Error != nil {
			return &response.GenericError{
				Message: "Unable to create campaign image",
				Err:     result.Error,
			}
		}

		campaignImages = append(campaignImages, campaignImage)
	}

	// Update the campaign object with campaignImages
	campaign.CampaignImages = campaignImages

	// Return the campaign as a JSON response
	return c.JSON(campaign)
}
