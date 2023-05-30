package model

type UserNoti struct {
	NotiObjectId *string     `json:"noti_object_id" gorm:"primaryKey;not null"`
	NotiObject   *NotiObject `json:"noti_object" gorm:"foreignKey:NotiObjectId;references:Id;not null"`
	UserId       *string     `json:"user_id" gorm:"primaryKey;not null"`
	User         *User       `json:"user" gorm:"foreignKey:UserId;not null"`
}
