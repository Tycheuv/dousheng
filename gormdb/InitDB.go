package gormdb

import (
	"dousheng/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"os/exec"
	"strings"
)

var DB *gorm.DB

func CreateTable() {
	sqlcmd, _ := os.ReadFile("./gormdb/table.sql")

	sqlArr := strings.Split(string(sqlcmd), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := DB.Exec(sql).Error
		if err != nil {
			log.Println("数据库导入失败:" + err.Error())
		} else {
			log.Println("success!")
		}
	}
}

func InitDB() error {
	command := `./gormdb/database.sh`
	err := os.Chmod(command, 0777) //提升文件权限
	if err != nil {
		return err
	}
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return err
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	config.GetDatabase()
	info := config.Info //
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		info.Name, info.Password, info.Host, info.Port, info.Database, info.Charset, info.ParseTime, info.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Mysql connect failed, err:%v", err)
		return err
	} else {
		fmt.Println("Mysql connect success!")
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		sqlDB.SetMaxIdleConns(info.MaxCon)   //设置空闲时的最大连接数
		sqlDB.SetMaxOpenConns(info.MaxOpCon) //设置与数据库的最大打开连接数
		sqlDB.SetConnMaxLifetime(-1)
		DB = db
		CreateTable()
		return nil
	}
}
