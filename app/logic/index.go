package logic

import (
	"fmt"
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

	result := models.GetVoteWithOptions(voteId)
	context.JSON(http.StatusOK, tools.ECode{
		Code:    0,
		Message: "get vote info successful",
		Data:    result,
	})
}

// DoVote 拿到复选框里的投票数据 写入数据库
func DoVote(context *gin.Context) {
	// 问题一 如何拿到投票用户的id	-- 登录写cookie + 中间件(登录状态才能投票)
	session := tools.GetSession(context)
	id := session["id"] //谁
	//将interface 转换为 int
	// 使用类型断言将 interface{} 类型转换为 int 类型
	var userID int
	var ok bool
	if userID, ok = id.(int); ok {
		// 类型断言成功，将值转换为 int 类型
		fmt.Println("The value is:", userID)
	} else {
		// 类型断言失败，原始值不是 int 类型
		fmt.Println("The value is not an int")
	}
	// session代替cookie？
	//id, _ := context.Cookie("id")
	// 如何从表单中哪信息
	voteId, _ := context.GetPostForm("vote_id")     // 哪个项目
	options, _ := context.GetPostFormArray("opt[]") // 哪个选项
	//处理前端 传回的数据
	voteID, _ := strconv.ParseInt(voteId, 10, 32)
	optionIDs := make([]int, 0)
	for _, opt := range options {
		optID, _ := strconv.ParseInt(opt, 10, 32)
		optionIDs = append(optionIDs, int(optID))
	}

	// 多选的数据以什么形式传回数据库 并 写入表中
	models.DoVote(userID, int(voteID), optionIDs)
	context.JSON(http.StatusOK, tools.OK)

}
