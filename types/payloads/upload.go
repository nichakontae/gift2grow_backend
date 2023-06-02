package payloads

type UploadCampaign struct {
	CampaignId *uint64 `form:"campaignId"`
}

type UploadProfile struct {
	UserId *string `form:"userId"`
}


type UploadCoverImg struct {
	Image *string `json:"image"`
}
