package payloads


type Tracking struct {
	CampaignId     *uint64    `json:"campaign_id"`
	UserId         *string    `json:"user_id"`
	TrackingNumber *string    `json:"tracking_number" `
}


