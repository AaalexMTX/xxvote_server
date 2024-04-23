package app

import (
	"fmt"
	"xxvote_server/app/config"
	"xxvote_server/app/models"
	"xxvote_server/app/routers"
)

// apiFox 做接口文档
// 规范使用log

func Start() {
	fmt.Println("start...")
	config.InitConfig()
	models.NewMysql()
	routers.NewIndexRouter()
}
