package payloads

type UserProfileForShare struct {
	Id           *string `json:"id"`
	Username     *string `json:"username"`
	ProfileImage *string `json:"profile_image"`
	TamboonPoint *int    `json:"tamboon_point"`
}
