package app

import "xxvote_server/app/routers"

func Start() {
	routers.NewIndexRouter()
}
