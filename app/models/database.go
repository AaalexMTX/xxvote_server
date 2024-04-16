package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConnect *gorm.DB

func NewMysql() {
	//config 后续用 viper来管理！！！
	address := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "v123456", "8.210.175.17:3306", "xxvote_server")
	//创建MYSQL 连接  -- 建立gorm.Open时传 gorm.Config是什么？？
	var mysqlConnect *gorm.DB
	var err error
	if mysqlConnect, err = gorm.Open(mysql.Open(address), &gorm.Config{}); err != nil {
		fmt.Printf("[NewMysql]mysql connection failure -- err:%s", err.Error())
		panic(err)
	}
	MysqlConnect = mysqlConnect

}
