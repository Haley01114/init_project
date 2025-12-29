package main

import (
	task "github.com/Haley01114/init_project/task3/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 博客系统
func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	// 建表
	task.CreateTable(db)

	// 初始化数据
	task.InitData(db)

	// 关联查询：1.1.查询一个用户：所有文章、评论信息
	task.QueryOneUser(db)
	// 关联查询：1.2.查询所有用户：所有文章、评论信息
	// task.QueryAllUser(db)
	// 关联查询：2.查询评论数量最多的文章信息
	task.QueryPostByCount(db)

	// 钩子函数：1.创建文章时，更新对应用户的文章数量字段 => AfterCreate
	// 钩子函数：2.删除评论时，如果评论数量为0，则更新文章的评论状态为“无评论”
	// 删除一个评论数量为1的评论，测试钩子函数2
	task.DeleteCommentByID(db)
}
