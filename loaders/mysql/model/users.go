package model

type User struct {
	Id           *string `json:"id" gorm:"primaryKey;not null"`
	Username     *string `json:"username" gorm:"not null"`
	ProfileImage *string `json:"profile_image" gorm:"not null"`
	FirstName    *string `json:"first_name" gorm:"not null"`
	LastName     *string `json:"last_name" gorm:"not null"`
	Email        *string `json:"email" gorm:"not null"`
	RankName     *string `json:"rank_name" gorm:"not null;index"`
	Rank         *Rank   `json:"rank" gorm:"foreignKey:RankName;references:RankName;not null"`
	TamboonPoint *int    `json:"tamboon_point" gorm:"not null"`
}
