package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote_server/app/models"
	"xxvote_server/app/tools"
)

func LoadingVotePage(context *gin.Context) {
	context.HTML(http.StatusOK, "votePage.html", gin.H{})
}

func LoadingRoot(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

// GetVotes 返回 所有vote记录（status > 0）
func GetVotes(context *gin.Context) {
	var votes = models.GetVotes()
	//响应带着 HTML页面、往页面里传gin.H数据（数据渲染该页面）
	context.JSON(http.StatusOK, tools.ECode{
		Data: votes,
	})
}
