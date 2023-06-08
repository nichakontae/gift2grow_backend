package ranking

import (
	"gift2grow_backend/loaders/mysql"
	"gift2grow_backend/loaders/mysql/model"
	"gift2grow_backend/types/payloads"
	"gift2grow_backend/types/response"
	"gift2grow_backend/utils/value"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	result := mysql.Gorm.Find(&users)
	if result.Error != nil {
		return &response.GenericError{
			Message: "Unable to get all users",
			Err:     result.Error,
		}
	}

	// Sort users by TamboonPoint in descending order
	sort.Slice(users, func(i, j int) bool {
		if users[i].TamboonPoint == nil || users[j].TamboonPoint == nil {
			return false // Handle nil values as needed
		}
		return *users[i].TamboonPoint > *users[j].TamboonPoint
	})

	data, _ := value.Iterate(users, func(user model.User) (*payloads.ShowRank, error) {
		// ProfileImage, _ := url.JoinPath(config.C.ProductionURL, *user.ProfileImage)
		return &payloads.ShowRank{
			UserId:        user.Id,
			Username:      user.Username,
			ProfileImage:  user.ProfileImage,
			Rank:          user.Rank,
			TamboonPoint:  user.TamboonPoint,
		}, nil
	})

	return c.JSON(response.NewResponse(data))
}


