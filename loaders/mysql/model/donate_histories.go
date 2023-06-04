package model

import "time"

type DonateHistory struct {
	TrackingNumber *string    `json:"tracking_number" gorm:"primaryKey;not null"`
	CampaignId     *uint64    `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign       *Campaign  `json:"campaign" gorm:"foreignKey:CampaignId;references:Id;not null"`
	UserId         *string    `json:"user_id" gorm:"primaryKey;not null"`
	User           *User      `json:"user" gorm:"foreignKey:UserId;references:Id;not null"`
	DonationDate   *time.Time `json:"donation_date" gorm:"not null"`
}
