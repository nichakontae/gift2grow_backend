package payloads

type RegisterAccount struct {
	UserId    *string `json:"user_id"`
	Username  *string `json:"username"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
}
