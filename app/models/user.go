package models

import "fmt"

// CreateUser 传user 数据库中创建用户记录
func CreateUser(user *User) bool {
	//Create 不直接返回err（return *DB） ，但是可以用*DB的Error属性，来判断是否连接成功
	if err := MysqlConnect.Create(user); err.Error != nil {
		fmt.Printf("[CreateUser]DB Create Fault -- err:%s\n", err.Error)
		return false
	}
	return true
}
