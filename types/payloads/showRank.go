package payloads

import "gift2grow_backend/types/enum"

type ShowRank struct {
	UserId       *string `json:"userId"`
	Username     *string `json:"username"`
	ProfileImage *string `json:"profile_image"`
	Rank         *enum.Rank `json:"rank"`
	TamboonPoint *int    `json:"tamboon_point"`
}