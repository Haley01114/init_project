package database

import (
	models2 "github.com/Haley01114/init_project/task4/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB 连接数据库
func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	user := models2.User{}
	post := models2.Post{}
	comment := models2.Comment{}
	AutoMigrateErr := db.AutoMigrate(&user, &post, &comment)
	if AutoMigrateErr != nil {
		return
	}

	DB = db
}
