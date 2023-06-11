package model

import "time"

type NotiObject struct {
	Id         *uint64     `json:"id" gorm:"primaryKey;not null"`
	CreatedAt  *time.Time  `json:"created_at" gorm:"not null"`
	Desc       *string     `json:"desc"`
	CampaignId *uint64     `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign   `json:"campaign" gorm:"foreignKey:CampaignId;references:Id;not null"`
	UserNoti   []*UserNoti `json:"user_noti" gorm:"foreignKey:NotiObjectId;references:Id;"`
}
