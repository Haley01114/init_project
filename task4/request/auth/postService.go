package auth

import (
	"net/http"

	"github.com/Haley01114/init_project/task4/database"
	"github.com/Haley01114/init_project/task4/database/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type createPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type getPostRequest struct {
	TitleID string `json:"titleId"`
}

type updatePostRequest struct {
	PostID  uint   `json:"postId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type deletePostRequest struct {
	PostID uint `json:"postId" binding:"required"`
}

// CreatePost 创建文章
func CreatePost(c *gin.Context) {
	// JWT token 匹配成功后，会在上下文中为 user_id,username赋值
	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
	}
	// 文章内容校验
	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章标题和文章内容不能为空"})
		c.Abort()
		return
	}
	// 保存到数据库
	post := &models.Post{
		UserID:  c.GetUint("user_id"),
		Title:   req.Title,
		Content: req.Content,
	}
	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}

	user := &models.User{}
	if err := database.DB.Debug().Model(&user).
		Where("id = ?", post.UserID).
		Update("posts_count", gorm.Expr("posts_count + 1")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章数量更新失败！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "文章创建成功~",
		"username": username,
		"title":    post.Title,
	})
}

// GetPostList 查询文章列表
func GetPostList(c *gin.Context) {
	// JWT token 匹配成功后，会在上下文中为 user_id,username赋值
	username := c.GetString("username")
	if username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
	}

	// 查询文章列表
	var req getPostRequest
	posts := &[]models.Post{}
	conditions := make(map[string]interface{})
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求错误"})
	}
	if req.TitleID != "" {
		conditions["title_id"] = req.TitleID
	}
	database.DB.Debug().Model(&posts).Where(conditions).Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"message": "查询文章列表成功~",
		"posts":   posts,
	})
}

// UpdatePost 文章更新
func UpdatePost(c *gin.Context) {
	var req updatePostRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		c.Abort()
		return
	}

	post := &models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  c.GetUint("user_id"),
	}
	if err := database.DB.Debug().Model(&post).
		Where("id = ?", req.PostID).
		Updates(post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章更新失败！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功~",
		"post":    post,
	})
}

// DeletePost 文章删除
func DeletePost(c *gin.Context) {
	var req deletePostRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		c.Abort()
		return
	}

	if err := database.DB.Debug().Delete(&models.Post{}, req.PostID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败！"})
		return
	}

	user := &models.User{}
	if err := database.DB.Debug().Model(&user).
		Where("id = ?", c.GetUint("user_id")).
		Update("posts_count", gorm.Expr("posts_count - 1")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章数量更新失败！"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "文章删除成功~",
		"postID":  req.PostID,
	})
}
