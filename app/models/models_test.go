package models

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
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
}
