package model

type WantList struct {
	CampaignId *uint64   `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign `json:"campaign" gorm:"foreignKey:CampaignId;references:Id;not null"`
	WantItem   *string   `json:"want_item" gorm:"primaryKey;not null"`
}
