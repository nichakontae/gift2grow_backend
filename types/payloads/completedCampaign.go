package payloads

type CompletedCampaign struct {
	CampaignId            *uint64   `json:"campaignId"`
	CoverImage            *string   `json:"coverImage"`
	Topic                 *string   `json:"topic"`
	SchoolName            *string   `json:"schoolName"`
	Description           *string   `json:"description"`
	IsCompleted           *bool     `json:"isCompleted"`
	CompletedAmount       *int      `json:"completedAmount"`
	EvidenceCampaignImage []string `json:"evidenceImg"`
	LetterOfThanks        *string   `json:"letterThank"`
}
