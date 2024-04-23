package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xxvote_server/app/models"
	"xxvote_server/app/tools"
)

/*
* trouble
* 表格无数据的情况下 发情查询数据详情（后台会报错！！！！（异常而非无数据））
 */

// ResultData 新定义返回结构 底层数据结构 -> 视图层结构
type ResultData struct {
	Title string
	Count int
	Opt   []*ResultVoteOpt
}
type ResultVoteOpt struct {
	Name  string
	Count int
}

func LoadingResultPage(context *gin.Context) {
	context.HTML(http.StatusOK, "result.html", gin.H{})
}

// GetVoteResult result——Ajax的异步请求--EChart展示
func GetVoteResult(context *gin.Context) {
	//拿到要查 投票目录的id
	id, _ := strconv.ParseInt(context.Query("id"), 10, 32)
	var voteWithOptions = models.GetVoteWithOptions(int(id))
	var result = ResultData{
		Title: voteWithOptions.Vote.Title,
		Count: 0,
		Opt:   nil,
	}
	for _, val := range voteWithOptions.Option {
		result.Count += val.Count
		temp := ResultVoteOpt{
			Name:  val.Name,
			Count: val.Count,
		}
		result.Opt = append(result.Opt, &temp)
	}
	// 返回一个JSON格式的数据
	context.JSON(http.StatusOK, tools.ECode{
		Data: result,
	})
}
