package model

import (
	"time"
)

type User struct {
	//
	Id int64 `gorm:"primaryKey" json:"id" `
	//
	Name string `json:"name"`
	//
	HeaderImg string `json:"header_img"`
	//
	Email string `json:"email"`
	//
	Mobile string `json:"mobile"`
	//
	IsOnline int64 `json:"is_online"`
	//
	CreatedAt time.Time `gorm:"autoCreateTime" `
	//
	UpdatedAt time.Time `gorm:"autoUpdateTime" `
}

func (u *User) TableName() string {
	return "c_user"
}
