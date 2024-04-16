package models

import (
	"fmt"
	"testing"
	"time"
	"xxvote_server/app/config"
)

func TestCreateUser(t *testing.T) {
	//Test 作为单例测试是 独立执行的 记得初始化配置！！
	config.InitConfig()
	NewMysql()
	var user = User{
		Uid:        11,
		Name:       "t",
		Password:   "123",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if ok := CreateUser(&user); !ok {
		fmt.Println("[CreateUser] Create fault")
	}
	CloseMysql()
}
