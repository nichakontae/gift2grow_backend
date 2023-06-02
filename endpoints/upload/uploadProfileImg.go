package upload

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

func ProfileImg(c *fiber.Ctx) error {

	body := new(payloads.UploadProfile)
	if err := c.BodyParser(body); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse image file",
			Err:     err,
		}
	}
	file, err := fileHeader.Open()
	if err != nil {
		return &response.GenericError{
			Message: "Unable to parse image file",
			Err:     err,
		}
	}

	// * Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		return &response.GenericError{
			Message: "Unable to decode image",
			Err:     err,
		}
	}

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

	var user model.User

	if result := mysql.Gorm.Model(&user).Where("id = ?", body.UserId).Update("profile_image", &fileName); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to fetch campaign image",
			Err:     result.Error,
		}
	}

	return c.JSON(&response.InfoResponse{
		Success: true,
		Message: "Updated image already",
	})

}
