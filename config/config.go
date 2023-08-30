package config

import (
	"os"
)

type DBInfo struct {
	Name      string
	Password  string
	Host      string
	Port      string
	Database  string
	Charset   string
	ParseTime string
	Loc       string
	MaxCon    int
	MaxOpCon  int
}

// Info 数据库信息配置
var Info = DBInfo{
	Name:      "root",      //用户名
	Password:  "123456",    //密码
	Host:      "127.0.0.1", //ip地址
	Port:      "3306",      //端口
	Database:  "dousheng",  //数据库名
	MaxCon:    30,          //最大连接数
	MaxOpCon:  3000,        //最大打开连接数
	Charset:   "utf8mb4",
	ParseTime: "True",
	Loc:       "Local",
}

func GetDatabase() {
	var ok bool
	var info = DBInfo{
		Database:  "dousheng", //数据
		MaxCon:    30,         //最大连接数
		MaxOpCon:  3000,       //最大打开连接数
		Charset:   "utf8mb4",
		ParseTime: "True",
		Loc:       "Local",
	}
	info.Name, ok = os.LookupEnv("MYSQL_USER")
	if !ok {
		return
	}
	info.Password, ok = os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		return
	}
	info.Host, ok = os.LookupEnv("MYSQL_HOST")
	if !ok {
		return
	}
	info.Port, ok = os.LookupEnv("MYSQL_PORT")
	if !ok {
		return
	}
	Info = info
	return
}

// VideoSavePath 项目运行网络ip地址或域名
var VideoSavePath = "" //例如 http://192.168.1.1:8080

// Host 项目运行端口
var Host = "8080"
