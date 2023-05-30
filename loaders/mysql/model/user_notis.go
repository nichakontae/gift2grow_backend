package model

type UserNoti struct {
	NotiObjectID *string     `json:"noti_object_id" gorm:"primaryKey;not null"`
	NotiObject   *NotiObject `json:"noti_object" gorm:"foreignKey:NotiObjectID;references:NotiObjectID"`
	UserID       *string     `json:"user_id" gorm:"primaryKey;not null"`
	User         *User       `json:"user" gorm:"foreignKey:UserID;references:UserID"`
}
