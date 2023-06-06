package model

type UserToken struct {
	UserId *string `json:"user_id" gorm:"primaryKey;not null"`
	User   *User   `json:"user" gorm:"foreignKey:UserId;references:Id;not null"`
	Token  *string `json:"token" gorm:"primaryKey;not null"`
}
