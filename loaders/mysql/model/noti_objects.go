package model

import "time"

type NotiObject struct {
	Id         *uint64     `json:"id" gorm:"primaryKey;not null"`
	IsRead     *bool       `json:"is_read" gorm:"not null"`
	CreatedAt  *time.Time  `json:"created_at" gorm:"not null"`
	CampaignId *uint64     `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign   `json:"campaign" gorm:"foreignKey:CampaignId;referencesId;not null"`
	UserNoti   []*UserNoti `json:"user_noti" gorm:"foreignKey:NotiObjectId;references:Id;"`
}
