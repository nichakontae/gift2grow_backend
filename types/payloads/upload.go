package payloads

type UploadCampaign struct {
	CampaignId *uint64 `form:"campaignId"`
}

type UploadCoverImg struct {
	Image *string `json:"image"`
}
