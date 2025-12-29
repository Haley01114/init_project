package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Remark string
	UserID uint
	User   User `gorm:"foreignkey:UserID"`
	PostID uint
	Post   Post `gorm:"foreignkey:PostID"`
}
