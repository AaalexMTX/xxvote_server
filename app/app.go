package app

import (
	"fmt"
	"xxvote_server/app/config"
	"xxvote_server/app/models"
	"xxvote_server/app/routers"
)

func Start() {
	fmt.Println("start...")
	config.InitConfig()
	models.NewMysql()
	routers.NewIndexRouter()
}
