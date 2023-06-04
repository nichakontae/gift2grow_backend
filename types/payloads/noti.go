package payloads

type EditNoti struct {
	NotiObjectId *uint64 `json:"notiObjectId"`
	UserId       *string `json:"userId"`
}

type PostNoti struct {
	CampaignId *uint64  `json:"campaignId"`
	UserIds    []string `json:"userIds"`
}
