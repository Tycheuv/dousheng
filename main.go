package main

import (
	"dousheng/config"
	"dousheng/gormdb"
	"github.com/gin-gonic/gin"
)

func main() {

	err := gormdb.InitDB() //连接数据库并建表
	if err != nil {
		return
	}
	config.GetUrl() //获取项目运行地址

	r := gin.Default() //获取默认路由

	initRouter(r) //初始化路由

	err = r.Run(":8080") //端口启动

	if err != nil {
		return
	}
}
