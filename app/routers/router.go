package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xxvote_server/app/logic"
)

func NewIndexRouter() {
	//设置默认根路由
	login := gin.Default()
	//设置html模板——是参数匹配模式（存在文件夹就要细分 *。html文件）
	login.LoadHTMLGlob("web/html/*.html")
	// 设置静态文件服务
	login.Static("/web", "./web")
	vote := login.Group("")
	vote.Use(checkUser)
	{
		//根路由
		login.GET("/index", logic.LoadingRoot)
		//login
		login.GET("/login", logic.GetLogin)
		login.POST("/login", logic.DoLogin)
		login.POST("/user/create", logic.CreateUser)

		//vote
		login.GET("/vote", logic.LoadingVotePage)
		login.POST("/do_vote", logic.DoVote)
		login.GET("/vote/infos", logic.GetVoteInfo)
		login.GET("/votes", logic.GetVotes)

		//result
		login.GET("/result", logic.LoadingResultPage)
		login.GET("/result/info", logic.GetVoteResult)
	}

	if err := login.Run(":8080"); err != nil {
		fmt.Printf("[NewIndexRouter]Failed to enable port monitoring -- err:%s\n", err.Error())
	}
}

func checkUser(context *gin.Context) {

}
