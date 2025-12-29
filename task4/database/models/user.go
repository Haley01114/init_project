package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Password   string
	Eamil      string
	PostsCount uint
	Posts      []Post `gorm:"foreignkey:UserID"`
}

// HashPassword 密码进行加密存储
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword 匹配加密后的password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
