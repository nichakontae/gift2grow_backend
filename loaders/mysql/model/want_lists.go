package model

type WantList struct {
	CampaignID *string   `json:"campaign_id" gorm:"primaryKey;not null"`
	Campaign   *Campaign `json:"campaign" gorm:"foreignKey:CampaignID;references:CampaignID"`
	WantItem   *string   `json:"want_item" gorm:"primaryKey;not null"`
}
