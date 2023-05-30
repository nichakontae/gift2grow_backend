package model

import "time"

type NotiObject struct {
	NotiObjectID *string    `json:"noti_object_id" gorm:"primaryKey;not null"`
	IsRead       *bool      `json:"is_read" gorm:"not null"`
	CreatedAt    *time.Time `json:"created_at" gorm:"not null"`
	Campaign     *Campaign  `json:"campaign" gorm:"foreignKey:CampaignID;references:CampaignID"`
}
