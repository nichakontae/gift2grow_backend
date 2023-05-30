package model

type Rank struct {
	RankName *string `json:"rank_name" gorm:"primaryKey;not null"`
	Slogan   *string `json:"slogan" gorm:"not null"`
	MinPoint *int    `json:"min_point" gorm:"not null"`
	MaxPoint *int    `json:"max_point" gorm:"not null"`
}
