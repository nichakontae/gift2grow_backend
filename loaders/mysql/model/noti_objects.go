package model

import "time"

type NotiObject struct {
	Id         *string    `json:"id" gorm:"primaryKey;not null"`
	IsRead     *bool      `json:"is_read" gorm:"not null"`
	CreatedAt  *time.Time `json:"created_at" gorm:"not null"`
	CampaignId *string    `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign  `json:"campaign" gorm:"foreignKey:CampaignId;referencesId;not null"`
}
