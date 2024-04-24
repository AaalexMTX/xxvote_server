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

	newUser := User{
		Uid:        111,
		Name:       "T2",
		Password:   "2222",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	// 执行的函数失败会报擦如错误 不用在Test里特意写
	if err := CreateUser(&newUser); err != nil {
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

func TestGetUserByName(t *testing.T) {
	config.InitConfig()
	NewMysql()
	var user = GetUserByName("admin")
	fmt.Printf("user: %+v", user)
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

func TestDoVote(t *testing.T) {
	config.InitConfig()
	NewMysql()
	if ok := DoVote(1, 2, []int{4}); ok == false {
		fmt.Println("vote failure")
	}
	fmt.Println("vote success")
}
