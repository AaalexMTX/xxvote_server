package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote_server/app/models"
	"xxvote_server/app/tools"
)

type loginUserInfo struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

//CaptchaId    string `json:"captcha_id" form:"captcha_id"`
//CaptchaValue string `json:"captcha_value" form:"captcha_value"`

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", gin.H{})
}

func DoLogin(context *gin.Context) {
	//将表单传入的 信息 绑在logic层的数据结构中
	var userInfo loginUserInfo

	if err := context.ShouldBind(&userInfo); err != nil {
		fmt.Println("[DoLogin]bind info fault")
	}

	//mysql 查询用户是否存在
	var user *models.User
	if user = models.GetUserStruct(userInfo.UserName, userInfo.Password); user.Id < 0 {
		context.JSON(http.StatusOK, tools.UserErr)
		return
	}

	//都验证成功 登录并设置session
	_ = tools.SetSession(context, userInfo.UserName, user.Id)

	//传消息和重定向页面无法同时进行
	context.JSON(http.StatusOK, tools.OK)
	context.Redirect(http.StatusSeeOther, "/vote")
}

func DoLogOut(context *gin.Context) {
	// 要加入session的清除
	// 重定向到login界面
	//context.HTML(http.StatusOK,"login.html",gin.H{})
	context.Redirect(http.StatusFound, "/login")
}
