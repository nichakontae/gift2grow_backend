package model

import "time"

type DonateHistory struct {
	CampaignID     *string    `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign       *Campaign  `json:"campaign" gorm:"foreignKey:CampaignID;references:CampaignID"`
	UserID         *string    `json:"user_id" gorm:"primaryKey;not null"`
	User           *User      `json:"user" gorm:"foreignKey:UserID;references:UserID"`
	TrackingNumber *string    `json:"tracking_number" gorm:"not null"`
	DonationDate   *time.Time `json:"donation_date" gorm:"not null"`
}
