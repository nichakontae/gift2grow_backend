package endpoints

import (
	"gift2grow_backend/endpoints/auth"
	"gift2grow_backend/endpoints/campaign_detail"
	"gift2grow_backend/endpoints/home"
	"gift2grow_backend/endpoints/notification"
	"gift2grow_backend/endpoints/profile"
	"gift2grow_backend/endpoints/ranking"
	"gift2grow_backend/endpoints/upload"

	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	authGroup := router.Group("/auth")
	authGroup.Get("/hello", auth.Hello)
	authGroup.Post("/register", auth.Register)

	campaignGroup := router.Group("/campaign")
	campaignGroup.Get("/getCampaignDetail", campaign_detail.GetCampaignDetail)
	campaignGroup.Post("/postTracking", campaign_detail.PostTracking)
	campaignGroup.Get("/completedCampaign", campaign_detail.GetCompletedCam)

	homeGroup := router.Group("/home")
	homeGroup.Get("/getAllCampaign", home.GetAllCampaign)

	notiGroup := router.Group("/noti")
	notiGroup.Get("/hello", notification.Hello)

	profileGroup := router.Group("/profile")
	profileGroup.Get("/hello", profile.Hello)

	rankGroup := router.Group("/rank")
	rankGroup.Get("/hello", ranking.Hello)

	uploadGroup := router.Group("/upload")
	uploadGroup.Post("/coverImg", upload.CoverImg)
	uploadGroup.Post("/campaignImg", upload.CampaignImg)
	uploadGroup.Post("/evidenceImg", upload.EvidenceImg)
	uploadGroup.Post("/profileImg", upload.ProfileImg)
}
