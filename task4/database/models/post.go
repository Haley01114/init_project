package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title         string
	Content       string
	CommentsCount uint
	CommentsState string
	UserID        uint
	User          User      `gorm:"foreignkey:UserID"`
	Comments      []Comment `gorm:"foreignkey:PostID"`
}
