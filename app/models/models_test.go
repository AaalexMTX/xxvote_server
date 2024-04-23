package models

import (
	"fmt"
	"testing"
	"time"
	"xxvote_server/app/config"
)

// user
func TestCreateUser(t *testing.T) {
	//Test 作为单例测试是 独立执行的 记得初始化配置！！
	config.InitConfig()
	NewMysql()
	defer CloseMysql()

	var user = User{
		Uid:        11,
		Name:       "t",
		Password:   "123",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	// 执行的函数失败会报擦如错误 不用在Test里特意写
	if ok := CreateUser(&user); !ok {
		fmt.Println("[TestCreateUser] Create fault")
		return
	}
	fmt.Println("[TestCreateUser] Create userRecord success")
}

func TestGetUser(t *testing.T) {
	config.InitConfig()
	NewMysql()
	defer CloseMysql()
	var userName = "admin"
	var password = "123"
	if ok := CheckUser(userName, password); ok == false {
		fmt.Println("[TestGetUser] user not exist")
		return
	}
	fmt.Println("[TestGetUser] user exist")
}

func TestGetUserStruct(t *testing.T) {
	config.InitConfig()
	NewMysql()
	defer CloseMysql()
	var userName = "admin"
	var password = "123"
	var user *User
	user = GetUserStruct(userName, password)
	fmt.Printf("%s\n%s\n", user.Name, user.Password)
}

// vote
func TestGetVotes(t *testing.T) {
	config.InitConfig()
	NewMysql()
	var votes = GetVotes()
	fmt.Printf("[TestGetVotes] votes :%+v", votes)
}

func TestGetVoteWithOptions(t *testing.T) {
	config.InitConfig()
	NewMysql()

	var ans = GetVoteWithOptions(1)
	fmt.Printf("%+v", ans)
}
