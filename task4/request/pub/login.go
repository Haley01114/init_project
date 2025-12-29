package pub

import (
	"errors"
	"net/http"

	"github.com/Haley01114/init_project/task4/database"
	"github.com/Haley01114/init_project/task4/database/models"
	"github.com/Haley01114/init_project/task4/jwt/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 登录
func Login(c *gin.Context) {
	// 验证必填项: username、password
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// 验证 用户名
	user := models.User{}
	if err := database.DB.Where("username = ?", request.Username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
		return
	}

	// 验证 密码
	if err := user.CheckPassword(request.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 生成 JWT
	token, err := utils.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成Token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "登录成功",
		"token":    token,
		"username": user.Username,
	})
}
