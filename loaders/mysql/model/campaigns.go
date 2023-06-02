package model

import "time"

type Campaign struct {
	Id                    *uint64                  `json:"id" gorm:"primaryKey;not null"`
	CoverImage            *string                  `json:"cover_image" gorm:"not null"`
	Topic                 *string                  `json:"topic" gorm:"not null"`
	SchoolName            *string                  `json:"school_name" gorm:"not null"`
	Location              *string                  `json:"location" gorm:"not null"`
	Description           *string                  `json:"description" gorm:"not null"`
	IsCompleted           *bool                    `json:"is_completed" gorm:"not null"`
	TelContact            *string                  `json:"tel_contact" gorm:"not null"`
	CompletedAmount       *int                     `json:"completed_amount" gorm:"not null"`
	CreatedAt             *time.Time               `json:"created_at" gorm:"not null"`
	UpdatedAt             *time.Time               `json:"updated_at" gorm:"not null"`
	LetterOfThanks        *string                  `json:"letter_of_thanks"`
	WantLists             []*WantList              `json:"want_lists" gorm:"foreignKey:CampaignId;references:Id;"`
	CampaignImages        []*CampaignImage         `json:"campaign_images" gorm:"foreignKey:CampaignId;references:Id;"`
	EvidenceCampaignImage []*EvidenceCampaignImage `json:"evidence_campaign_image" gorm:"foreignKey:CampaignId;references:Id;"`
}
