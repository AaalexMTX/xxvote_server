package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewIndexRouter() {
	//设置默认根路由
	index := gin.Default()
	//设置html模板——是参数匹配模式（存在文件夹就要细分 *。html文件）
	index.LoadHTMLGlob("web/html/*.html")
	// 设置静态文件服务
	index.Static("/web", "./web")
	{
		//根路由
		index.GET("/index", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", gin.H{})
		})
		index.GET("/login", func(context *gin.Context) {
			context.HTML(http.StatusOK, "login.html", gin.H{})
		})
	}

	if err := index.Run(":8080"); err != nil {
		fmt.Printf("[NewIndexRouter]Failed to enable port monitoring -- err:%s\n", err.Error())
	}
}
