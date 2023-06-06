package ranking

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"github.com/gofiber/fiber/v2"
)

func GetProfileForShare(c *fiber.Ctx) error {
	userId := c.Query("userId")
	if userId == "" {
		return &response.GenericError{
			Message: "User id not found",
			Err:     nil,
		}
	}
	var user model.User

	if result := mysql.Gorm.Where("id = ?", userId).First(&user); result.Error != nil {
		return &response.GenericError{
			Message: "User not found",
			Err:     nil,
		}
	}

	userProfileForShare := &payloads.UserProfileForShare{
		Id:           user.Id,
		ProfileImage: user.ProfileImage,
		Username:     user.Username,
		TamboonPoint: user.TamboonPoint,
	}

	return c.JSON(response.NewResponse(userProfileForShare))
}
