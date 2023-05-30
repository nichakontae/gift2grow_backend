package model

import "time"

type CampaignImage struct {
	CampaignId *string    `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign  `json:"campaign" gorm:"foreignKey:CampaignId;references:Id;not null"`
	Image      *string    `json:"image" gorm:"primaryKey;not null"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"not null"`
}
