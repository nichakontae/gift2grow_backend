package payloads

import (
	"time"
)

type AllCampaign struct {
    CampaignId        *uint64       `json:"campaignId"`
	CoverImage        *string       `json:"coverImage"`
	SchoolName        *string       `json:"schoolName"`
	IsCompleted       *bool         `json:"isCompleted"`
	CompletedAmount   *int          `json:"completedAmount"`
	TrackingAmount    *int64          `json:"trackingAmount"`
	CreatedAt         *time.Time    `json:"createdAt"`
}