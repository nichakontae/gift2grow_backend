package ranking

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/config"
	"gift2grow_backend/utils/value"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	if result := mysql.Gorm.Find(&users); result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all users",
			Err:     result.Error,
		}
	}

	data, _ := value.Iterate(users, func(user model.User) (*payloads.ShowRank, error) {
		ProfileImage, _ := url.JoinPath(config.C.ProductionURL, *user.ProfileImage)
		return &payloads.ShowRank{
			UserId:      user.Id,
			Username:    user.Username,
			ProfileImage:  &ProfileImage,
			//ProfileImage:  user.ProfileImage,
			Rank:    user.Rank,
			TamboonPoint:    user.TamboonPoint,
		}, nil
	})
	
	return c.JSON(data)
}
