package model

type WantList struct {
	CampaignId *string   `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign `json:"campaign" gorm:"foreignKey:CampaignId;references:Id;not null"`
	WantItem   *string   `json:"want_item" gorm:"primaryKey;not null"`
}
