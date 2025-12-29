package auth

import (
	"net/http"

	"github.com/Haley01114/init_project/task4/database"
	"github.com/Haley01114/init_project/task4/database/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createCommentRequest struct {
	Remark string `json:"remark" binding:"required"`
	PostID uint   `json:"postId" binding:"required"`
}

type getCommentListRequest struct {
	PostID uint `json:"postId" binding:"required"`
}

// CreateComment 发表评论
func CreateComment(c *gin.Context) {
	var req createCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		c.Abort()
		return
	}

	comment := &models.Comment{
		Remark: req.Remark,
		PostID: req.PostID,
		UserID: c.GetUint("user_id"),
	}

	if err := database.DB.Debug().Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论保存失败！"})
		c.Abort()
		return
	}

	post := &models.Post{}
	if err := database.DB.Model(&post).
		Where("id = ?", comment.PostID).
		Update("comments_count", gorm.Expr("comments_count + 1")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论数量更新失败！"})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "评论成功~",
		"data":    comment,
	})
}

// GetCommentList 查询评论列表
func GetCommentList(c *gin.Context) {
	// 查询文章列表
	var req getCommentListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误！"})
		c.Abort()
		return
	}

	var comments []models.Comment
	database.DB.Debug().Where("post_id = ?", req.PostID).Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"message":  "查询评论列表成功~",
		"comments": comments,
	})
}
