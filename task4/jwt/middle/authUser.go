package middle

import (
	"net/http"

	"github.com/Haley01114/init_project/task4/database"
	"github.com/Haley01114/init_project/task4/database/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type postRequest struct {
	PostID uint `json:"postId" binding:"required"`
}

// AuthUser 验证 文章 是否属于 当前用户
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postRequest
		if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "文章id 不能为空"})
			c.Abort()
			return
		}

		// 验证 文章 是否属于 当前用户
		post := models.Post{}
		database.DB.Debug().Where("id = ?", req.PostID).Find(&post)
		if post.UserID != c.GetUint("user_id") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户无权限！"})
			c.Abort()
			return
		}
		c.Next()
	}
}
