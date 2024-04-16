package app

import (
	"fmt"
	"xxvote_server/app/routers"
)

func Start() {
	fmt.Println("start...")
	routers.NewIndexRouter()
}
