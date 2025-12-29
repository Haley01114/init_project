package pub

import (
	"net/http"

	"github.com/Haley01114/init_project/task4/database"
	"github.com/Haley01114/init_project/task4/database/models"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Username: req.Username,
		Eamil:    req.Email,
	}

	// 加密密码
	if err := user.HashPassword(req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	// 保存到数据库
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "注册成功",
		"username": user.Username,
		"eamil":    user.Eamil,
	})
}
