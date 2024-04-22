package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xxvote_server/app/models"
	"xxvote_server/app/tools"
)

func LoadingVotePage(context *gin.Context) {
	context.HTML(http.StatusOK, "vote.html", gin.H{})
}

func LoadingRoot(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

// GetVotes 返回 所有vote记录（status > 0）
func GetVotes(context *gin.Context) {
	var votes = models.GetVotes()
	//响应带着 HTML页面、往页面里传gin.H数据（数据渲染该页面）
	context.JSON(http.StatusOK, tools.ECode{
		Data:    votes,
		Message: "Voting data request successful",
	})
}

func GetVoteInfo(context *gin.Context) {
	//从 Ajax 请求中拿到 申请的 voteId
	voteId, _ := strconv.Atoi(context.Query("id"))

	result := models.GetVoteWithOptions(int32(voteId))
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "get vote info successful",
		Data:    result,
	})
}

// DoVote 拿到复选框里的投票数据 写入数据库
func DoVote(context *gin.Context) {
	// 问题一 如何拿到投票用户的id	-- 登录写cookie + 中间件(登录状态才能投票)
	// 如何从表单中哪信息
	// 多选的数据以什么形式传回数据库 并 写入表中
}
