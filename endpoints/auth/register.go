package auth

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/enum"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	registerAccount := new(payloads.RegisterAccount)
	if err := c.BodyParser(registerAccount); err != nil {
		return &response.GenericError{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	rank := enum.LittleDandelion
	tamboonPoint := 0
	profileImage := "/images/"

	user := model.User{
		Id:           registerAccount.UserId,
		Username:     registerAccount.Username,
		ProfileImage: &profileImage,
		FirstName:    registerAccount.Firstname,
		LastName:     registerAccount.Lastname,
		Email:        registerAccount.Email,
		Rank:         &rank,
		TamboonPoint: &tamboonPoint,
	}
	if result := mysql.Gorm.Create(user); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to create an account",
			Err:     result.Error,
		}
	}
	return c.JSON(&response.InfoResponse{
		Success: true,
		Message: "Created an account already",
	})
}
