package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote_server/app/logic"
)

func NewIndexRouter() {
	//设置默认根路由
	login := gin.Default()
	//设置html模板——是参数匹配模式（存在文件夹就要细分 *。html文件）
	login.LoadHTMLGlob("web/html/*.html")
	// 设置静态文件服务
	login.Static("/web", "./web")
	{
		//根路由
		login.GET("/index", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", gin.H{})
		})
		login.GET("/login", logic.GetLogin)
		login.POST("/login", logic.DoLogin)
	}

	if err := login.Run(":8080"); err != nil {
		fmt.Printf("[NewIndexRouter]Failed to enable port monitoring -- err:%s\n", err.Error())
	}
}
