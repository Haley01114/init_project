package task3

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	PostsCount uint
	Posts      []Post `gorm:"foreignkey:UserID"`
}

type Post struct {
	gorm.Model
	Content       string
	CommentsCount uint
	CommentsState string
	UserID        uint
	User          User      `gorm:"foreignkey:UserID"`
	Comments      []Comment `gorm:"foreignkey:PostID"`
}

type Comment struct {
	gorm.Model
	Remark string
	UserID uint
	User   User `gorm:"foreignkey:UserID"`
	PostID uint
	Post   Post `gorm:"foreignkey:PostID"`
}

// 博客系统
// 建表
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

// 初始化数据
func InitData(db *gorm.DB) {
	users := []User{
		{Name: "张三", PostsCount: 2},
		{Name: "李四", PostsCount: 3},
	}
	db.Create(&users)
	posts := []Post{
		{Content: "张三文章1", CommentsCount: 2, CommentsState: "正常", UserID: users[0].ID},
		{Content: "张三文章2", CommentsCount: 1, CommentsState: "正常", UserID: users[0].ID},

		{Content: "李四文章1", CommentsCount: 1, CommentsState: "正常", UserID: users[1].ID},
		{Content: "李四文章2", CommentsCount: 3, CommentsState: "正常", UserID: users[1].ID},
		{Content: "李四文章3", CommentsCount: 0, CommentsState: "无评论", UserID: users[1].ID},
	}
	db.Create(&posts)
	comments := []Comment{
		{Remark: "张三文章1 评论1", UserID: users[0].ID, PostID: posts[0].ID},
		{Remark: "张三文章1 评论2", UserID: users[0].ID, PostID: posts[0].ID},
		{Remark: "张三文章2 评论1", UserID: users[0].ID, PostID: posts[1].ID},

		{Remark: "李四文章1 评论1", UserID: users[1].ID, PostID: posts[2].ID},
		{Remark: "李四文章2 评论1", UserID: users[1].ID, PostID: posts[3].ID},
		{Remark: "李四文章2 评论2", UserID: users[1].ID, PostID: posts[3].ID},
		{Remark: "李四文章2 评论3", UserID: users[1].ID, PostID: posts[3].ID},
	}
	db.Create(&comments)
}

// 查询一个用户：所有文章、评论信息
func QueryOneUser(db *gorm.DB) {
	var user User
	db.Debug().Preload("Posts.Comments").Find(&user, 1)
	fmt.Printf("用户: %s\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("  文章: %s\n", post.Content)
		for _, comment := range post.Comments {
			fmt.Printf("    评论: %s\n", comment.Remark)
		}
	}
}

// 查询所有用户：所有文章、评论信息
func QueryAllUser(db *gorm.DB) {
	var users []User
	db.Debug().Preload("Posts.Comments").Find(&users)
	for _, user := range users {
		fmt.Printf("用户: %s\n", user.Name)
		for _, post := range user.Posts {
			fmt.Printf("  文章: %s\n", post.Content)
			for _, comment := range post.Comments {
				fmt.Printf("    评论: %s\n", comment.Remark)
			}
		}
	}
}

// 查询 评论数量最多 的文章
func QueryPostByCount(db *gorm.DB) {
	var post Post
	db.Debug().Model(&Post{}).
		Select("posts.*,COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post)
	fmt.Printf("评论数最多的文章信息：\n%v\n", post)
}

// 钩子函数：1.创建文章时，更新对应用户的文章数量字段
// func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
// 	tx.Model(&User{}).
// 		Where("id = ?", p.UserID).
// 		Update("posts_count", gorm.Expr("posts_count + 1"))
// 	return
// }

// 钩子函数：2.删除评论时，如果评论数量为0，则更新文章的评论状态为“无评论”
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 如果评论数量 > 0，就将先将评论数量减一
	tx.Model(&Post{}).
		Where("id = ? AND comments_count > 0", c.PostID).
		Update("comments_count", gorm.Expr("comments_count - 1"))

	var post Post
	tx.Find(&post, c.PostID)
	if post.CommentsCount == 0 {
		tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comments_state", "无评论")
	}
	return
}

// 删除一个评论数量为1的评论，测试钩子函数2
func DeleteCommentByID(db *gorm.DB) {
	// 先查询一下，不然Comment中没有PostID，影响钩子函数
	var comment Comment
	db.Debug().First(&comment, 4)
	db.Debug().Delete(&comment)
}
