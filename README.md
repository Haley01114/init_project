go 语言学习：任务代码

目录:

[task1](task1)：Go语言基础

[task2](task2)：Go高级进阶

[task3](task3)：ORM框架：gorm

[task4](task4)：Web框架：gin/go-zero    
1.主程序： [main.go](task4/main.go) 

    端口：8080  
2.数据库设计及模型定义：[database](task4/database)  
    
    用户表：users
    文章表：posts
    评论表：comments
3.用户认证及授权：[authJWT.go](task4/jwt/middle/authJWT.go)   
    
    注册：注册成功后，将用户保存到users表
        curl --request POST \
        --url http://localhost:8080/user/register \
        --header 'Accept: */*' \
        --header 'Accept-Encoding: gzip, deflate, br' \
        --header 'Connection: keep-alive' \
        --header 'Content-Type: application/json' \
        --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
        --data '{
        "username":"用户2",
        "password":"123456",
        "email":"email@111111"
        }'

    登录：登录成功后生成 JWT Token，后续功能页面需要携带 JWT Token 访问
        curl --request POST \
        --url http://localhost:8080/user/login \
        --header 'Accept: */*' \
        --header 'Accept-Encoding: gzip, deflate, br' \
        --header 'Connection: keep-alive' \
        --header 'Content-Type: application/json' \
        --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
        --data '{
        "username":"用户1",
        "password":"123456"
        }'
4.文章管理：[postService.go](task4/request/auth/postService.go) 【校验 JWT Token】

    发表文章：保存到posts表，同时更新users表用户的文章数量
            curl --request POST \
            --url http://localhost:8080/api/createPost \
            --header 'Accept: */*' \
            --header 'Accept-Encoding: gzip, deflate, br' \
            --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6IueUqOaItzIiLCJzdWIiOiIyIiwiZXhwIjoxNzY3MDE1MDUxLCJpYXQiOjE3NjcwMTQ0NTF9.i0Hm2vuSwprDQ90Q8ucq_EjWrK6C58JvzHZ04bWP3Dw' \
            --header 'Connection: keep-alive' \
            --header 'Content-Type: application/json' \
            --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
            --data '{
            "title":"用户2-文章1",
            "content":"用户2-文章1-内容xxxxxx"
            }'
    读取文章：无参访问时查询所有文章列表，入参titleID时查某文章详情
            curl --request GET \
            --url http://localhost:8080/api/getPostList \
            --header 'Accept: */*' \
            --header 'Accept-Encoding: gzip, deflate, br' \
            --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IueUqOaItzEiLCJzdWIiOiIxIiwiZXhwIjoxNzY3MDE1MjIwLCJpYXQiOjE3NjcwMTQ2MjB9.5XaFE1WJBYjuqljwAMRpbULMRNtHnk0vmXJaXYHRtYk' \
            --header 'Connection: keep-alive' \
            --header 'Content-Type: application/json' \
            --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
            --data '{
            "titleId1":"用户1-文章1"
            }'
            
    文章更新：根据入参进行更新文章【校验当前用户是否是文章作者】
            curl --request POST \
            --url http://localhost:8080/api/post/updatePost \
            --header 'Accept: */*' \
            --header 'Accept-Encoding: gzip, deflate, br' \
            --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IueUqOaItzEiLCJzdWIiOiIxIiwiZXhwIjoxNzY3MDE3NTc1LCJpYXQiOjE3NjcwMTY5NzV9.QmuqMODPPMECSjF3aYsNafpK2l2HSyQ_nxMFwWUf5Qo' \
            --header 'Connection: keep-alive' \
            --header 'Content-Type: application/json' \
            --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
            --data '{
            "postId": 9,
            "title": "用户2-文章1",
            "content": "用户2-文章1-内容xx1111222"
            }'
            
    文章删除：根据id删除文章，并更新users表用户的文章数量【校验当前用户是否是文章作者】
            curl --request POST \
            --url http://localhost:8080/api/post/deletePost \
            --header 'Accept: */*' \
            --header 'Accept-Encoding: gzip, deflate, br' \
            --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IueUqOaItzEiLCJzdWIiOiIxIiwiZXhwIjoxNzY3MDE3OTg0LCJpYXQiOjE3NjcwMTczODR9.bNojdq31urBYP_JxXg6cfN2nA_WOLMQicc6s_Dx2neo' \
            --header 'Connection: keep-alive' \
            --header 'Content-Type: application/json' \
            --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
            --data '{
            "postId": 2
            }'

5.评论功能：[commentService.go](task4/request/auth/commentService.go) 【校验 JWT Token】

    发表评论：保存到comments表，同时更新posts表文章的评论数量
            curl --request POST \
            --url http://localhost:8080/api/createComment \
            --header 'Accept: */*' \
            --header 'Accept-Encoding: gzip, deflate, br' \
            --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IueUqOaItzEiLCJzdWIiOiIxIiwiZXhwIjoxNzY3MDE5Mjc5LCJpYXQiOjE3NjcwMTg2Nzl9.luPAlPZr7Y1uyT8eRCyHDSex60gA0ZeTDYr48KU6PUw' \
            --header 'Connection: keep-alive' \
            --header 'Content-Type: application/json' \
            --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
            --data '{
            "remark":"pinglun11111",
            "postId": 1
            }'

    查询评论：入参titleID时查某文章评论列表
            curl --request GET \
            --url http://localhost:8080/api/getCommentList \
            --header 'Accept: */*' \
            --header 'Accept-Encoding: gzip, deflate, br' \
            --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IueUqOaItzEiLCJzdWIiOiIxIiwiZXhwIjoxNzY3MDE5ODQ5LCJpYXQiOjE3NjcwMTkyNDl9.nv8zRB5OMYvirYCM_tXZXwRBjD_ZINTuqJdg7TzPbCI' \
            --header 'Connection: keep-alive' \
            --header 'Content-Type: application/json' \
            --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
            --data '{
            "postId":1
            }'

