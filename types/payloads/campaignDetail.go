package payloads

// import "time"

type Campaign struct {
	CampaignId            *uint64                  `json:"campaignId"`
	CoverImage            *string                  `json:"cover_image"`
	Topic                 *string                  `json:"topic"`
	Location              *string                  `json:"location"`
	SchoolName            *string                  `json:"school_name"`
	Description           *string                  `json:"description"`
	IsCompleted           *bool                    `json:"is_completed"`
	TelContact            *string                  `json:"tel_contact"`
	CompletedAmount       *int                     `json:"completed_amount"`
	CreatedAt             *string              `json:"created_at" `
	WantLists             []string             		`json:"want_lists" `
	CampaignImages        []string        			`json:"campaign_images"`
}
