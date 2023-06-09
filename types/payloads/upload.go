package payloads

type UploadCampaign struct {
	CampaignId *uint64 `form:"campaignId"`
}

type UploadCampaignEvidenceThank struct {
	CampaignId *uint64 `form:"campaignId"`
	LetterOfThanks *string `form:"letterThanks"`
}

type UploadProfile struct {
	UserId *string `form:"userId"`
}


type UploadCoverImg struct {
	Image *string `json:"image"`
}
