# 项目文档

### 环境需求

1. Go版本：1.20及以上
2. 数据库：MySQL8.0

### 使用方法

1. 克隆项目

```text
git clone https://github.com/Uvnams/dousheng.git
```

2. 添加依赖

```go
go mod tidy
```

3. 创建数据库

```text
Mysql>create database dousheng;
*本地运行需自行配置config/config.go中数据库信息
```

4. 运行

```go
go run main.go router.go
```

### 实现功能

实现了接口文档中给出的所有接口

+ 用户模块：注册、登录、获取用户信息
+ 视频流模块：发布视频、获取`Feed`流、查看个人已发布视频
+ 关注模块：关注操作、获取关注列表、获取粉丝列表
+ 评论模块：评论操作、获取评论列表
+ 点赞模块：点赞操作、获取点赞列表
+ 聊天模块：发送消息、获取对方消息

### 项目结构

```text
dousheng
|		main.go //项目入口
|		router.go //初始化路由
|		README
|
|--config
|		config.go //配置文件
|
|--controller
|		comment.go //评论接口	
|		common.go //结构体的定义
|		favorite.go //喜欢操作接口
|		feed.go //视频流接口
|		message.go //聊天接口
|		publish.go //视频发布、发布列表
|		relation.go //关注、好友接口
|		user.go //登陆、注册接口
|
|--gormdb
|		InitDB.go //Mysql初始化
|		table.sql //数据库表
|
└─public //发布的视频存储在此
```
