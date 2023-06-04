package payloads

import "gift2grow_backend/types/enum"

type User struct {
	Id            *string          `json:"id" gorm:"primaryKey;not null"`
	Username      *string          `json:"username" gorm:"not null"`
	ProfileImage  *string          `json:"profile_image" gorm:"not null"`
	FirstName     *string          `json:"first_name" gorm:"not null"`
	LastName      *string          `json:"last_name" gorm:"not null"`
	Email         *string          `json:"email" gorm:"not null"`
	Rank          *enum.Rank       `json:"rank" gorm:"not null;type:ENUM('Born to be Angle','Trainee Angel','Little Dandelion')"`
	TamboonPoint  *int             `json:"tamboon_point" gorm:"not null"`
	DonateHistory []*DonateHistory `json:"donate_history" gorm:"foreignKey:UserId;references:Id;"`
	UserNoti      []*UserNoti      `json:"user_noti" gorm:"foreignKey:UserId;references:Id;"`
}
