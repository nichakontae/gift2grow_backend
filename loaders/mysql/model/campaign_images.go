package model

import "time"

type CampaignImage struct {
	CampaignID *string    `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign  `json:"campaign" gorm:"foreignKey:CampaignID;references:CampaignID"`
	Image      *string    `json:"image" gorm:"primaryKey;not null"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"not null"`
}
