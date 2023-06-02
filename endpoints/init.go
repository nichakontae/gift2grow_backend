package endpoints

import (
	"gift2grow_backend/endpoints/auth"
	"gift2grow_backend/endpoints/campaign_detail"
	"gift2grow_backend/endpoints/home"
	"gift2grow_backend/endpoints/notification"
	"gift2grow_backend/endpoints/profile"
	"gift2grow_backend/endpoints/ranking"
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	authGroup := router.Group("/auth")
	authGroup.Get("/hello", auth.Hello)

	campaignGroup := router.Group("/campaign")
	campaignGroup.Get("/hello", campaign_detail.Hello)

	homeGroup := router.Group("/home")
	homeGroup.Get("/getAllCampaign", home.Hello)

	notiGroup := router.Group("/noti")
	notiGroup.Get("/hello", notification.Hello)

	profileGroup := router.Group("/profile")
	profileGroup.Get("/hello", profile.Hello)

	rankGroup := router.Group("/rank")
	rankGroup.Get("/hello", ranking.Hello)
}
