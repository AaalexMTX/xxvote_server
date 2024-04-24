package models

import "fmt"

// CreateUser 传user 数据库中创建用户记录
func CreateUser(user *User) error {
	//Create 不直接返回err（return *DB） ，但是可以用*DB的Error属性，来判断是否连接成功
	if err := MysqlConnect.Create(user); err != nil {
		fmt.Printf("[CreateUser]DB Create Fault -- err:%s\n", err.Error)
		return err.Error
	}
	return nil
}

// CheckUser 用唯一字段name 查用户记录
func CheckUser(userName string, password string) bool {
	//稍作修改   不把 model层的User结构向上暴露
	//校验 和 setSession在models层完成
	var user User
	if err := MysqlConnect.Raw("select * from user where name = ? limit 1", userName).Scan(&user).Error; err != nil {
		fmt.Printf("[CheckUser]query user not exist -- err:%s\n", err.Error())
		return false
	}
	//身份信息不对
	if password != user.Password {
		fmt.Printf("[CheckUser]password not correct\n")
		return false
	}
	return true
}

// GetUserStruct 难用的原因在于 冗杂了 查询和验证的功能！！！！
func GetUserStruct(userName string, password string) *User {
	//校验 和 setSession在models层完成
	var user User
	if err := MysqlConnect.Raw("select * from user where name = ? limit 1", userName).Scan(&user).Error; err != nil {
		fmt.Printf("[GetUserStruct]query user not exist -- err:%s\n", err.Error())
	}
	//身份信息不对
	if password != user.Password {
		fmt.Printf("[GetUserstruct]password not correct\n")
		return &user
	}
	return &user
}

func GetUserByName(userName string) *User {
	var user User
	// Scan(&user).Error 这里不加.Error 只表示是否扫描成功
	if err := MysqlConnect.Raw("select * from user where name = ?", userName).Scan(&user).Error; err != nil {
		fmt.Printf("[GetUserByName]query user not exist -- err:%s\n", err)
	}
	return &user
}
