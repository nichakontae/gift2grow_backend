package upload

import (
	"fmt"
	"gift2grow_backend/endpoints/notification"
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/config"
	"gift2grow_backend/utils/text"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func EvidenceImg(c *fiber.Ctx) error {
	// * Parse user JWT token
	//token := c.Locals("user").(*jwt.Token)
	//claims := token.Claims.(*common.UserClaim)

	// * Parse body
	body := new(payloads.UploadCampaignEvidenceThank)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	// * Parse multipart file parameter
	form, _ := c.MultipartForm()
	files := form.File["image"]
	for _, file := range files {
		fileHeader, err := file.Open()
		if err != nil {
			return &response.GenericError{
				Message: "Cannot open image",
				Err:     err,
			}
		}
		defer fileHeader.Close()

		// * Decode image
		img, _, err := image.Decode(fileHeader)
		if err != nil {
			return &response.GenericError{
				Message: "unable to decode file as image",
				Err:     err,
			}
		}

		// * Assign file path
		fileSalt := *text.GenerateString(text.GenerateStringSet.Num, 6)

		// * Save image to file
		savingFile, err := os.Create("./images/" + fileSalt + ".jpeg")
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
		// * Update user record
		evidenceImage := &model.EvidenceCampaignImage{
			CampaignId: body.CampaignId,
			Campaign:   nil,
			Image:      &fileName,
			UpdatedAt:  nil,
		}
		if result := mysql.Gorm.Create(evidenceImage); result.Error != nil {
			return &response.GenericError{
				Message: "Unable to fetch evidence of campaign image",
				Err:     result.Error,
			}
		}
	}

	var campaign model.Campaign

	if result := mysql.Gorm.Model(&campaign).Where("id = ?", body.CampaignId).Update("letter_of_thanks", body.LetterOfThanks); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to fetch campaign image",
			Err:     result.Error,
		}
	}
	c.Request().Header.Set("Content-Type", "application/json")
	c.Request().Header.Set("Authorization", config.C.AuthKey)

	campaignIDString := strconv.FormatUint(*body.CampaignId, 10)
	notification.NotifyUser(campaignIDString)

	return c.JSON(&response.InfoResponse{
		Success: true,
		Message: "Updated image already",
	})
}
