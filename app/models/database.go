package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlConnect *gorm.DB

func NewMysql() {
	//config 用 viper来管理！！！
	var mysqlUser = viper.GetString("mysql.user")
	var mysqlPassword = viper.GetString("mysql.password")
	var mysqlSocket = fmt.Sprintf("%s:%s", viper.GetString("mysql.ip"), viper.GetString("mysql.port"))
	var mysqlTable = viper.GetString("mysql.tableName")
	//fmt.Printf("%s\n%s\n%s\n%s\n", mysqlUser, mysqlPassword, mysqlSocket, mysqlTable)
	address := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlSocket, mysqlTable)

	//创建MYSQL 连接  -- 建立gorm.Open时传 gorm.Config是什么？？
	var mysqlConnect *gorm.DB
	var err error
	if mysqlConnect, err = gorm.Open(mysql.Open(address), &gorm.Config{}); err != nil {
		fmt.Println(address)
		panic(fmt.Sprintf("[NewMysql]mysql connection failure -- err:%s\n", err.Error()))
	}
	MysqlConnect = mysqlConnect
}

func CloseMysql() {
	c, _ := MysqlConnect.DB()
	_ = c.Close()
}
