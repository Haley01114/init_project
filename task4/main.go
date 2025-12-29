package main

import (
	"github.com/Haley01114/init_project/task4/database"
	"github.com/Haley01114/init_project/task4/jwt/middle"
	"github.com/Haley01114/init_project/task4/request/auth"
	"github.com/Haley01114/init_project/task4/request/pub"
	"github.com/gin-gonic/gin"
)

func main() {
	// 连接数据库
	database.ConnectDB()

	// gin 初始化
	router := gin.Default()

	// user 公共路由
	pubGroup := router.Group("/user")
	{
		// 注册
		pubGroup.POST("/register", pub.Register)
		// 登录
		pubGroup.POST("/login", pub.Login)
	}

	// 需要JWT认证 的路由
	apiGroup := router.Group("/api")
	apiGroup.Use(middle.AuthJWT())
	{
		// 发表文章
		apiGroup.POST("/createPost", auth.CreatePost)
		// 文章列表
		apiGroup.GET("/getPostList", auth.GetPostList)
		// 发表评论
		apiGroup.POST("/createComment", auth.CreateComment)
		// 发表列表
		apiGroup.GET("/getCommentList", auth.GetCommentList)
	}
	// 只当前用户可操作
	authUser := apiGroup.Group("/post")
	authUser.Use(middle.AuthUser())
	{
		// 更新文章
		authUser.POST("/updatePost", auth.UpdatePost)
		// 删除文章
		authUser.POST("/deletePost", auth.DeletePost)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
