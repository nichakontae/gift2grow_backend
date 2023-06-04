package payloads

import "time"

type EvidenceCampaignImage struct {
	CampaignId *uint64    `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign  `json:"campaign" gorm:"foreignKey:CampaignId;referencesId;not null"`
	Image      *string    `json:"image" gorm:"primaryKey;not null"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"not null"`
}
