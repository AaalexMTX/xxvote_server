package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewIndexRouter() {
	//设置默认根路由
	index := gin.Default()
	index.LoadHTMLGlob("web/html/*")

	{
		//根路由
		index.GET("/index", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}

	if err := index.Run(":8080"); err != nil {
		fmt.Printf("[NewIndexRouter]Failed to enable port monitoring -- err:%s\n", err.Error())
	}
}
