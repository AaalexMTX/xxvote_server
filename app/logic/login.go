package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xxvote_server/app/models"
)

type loginUserInfo struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	//CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	//CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

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
	if ok := models.GetUser(userInfo.UserName, userInfo.Password); ok == false {
		fmt.Printf("[DoLogin] info fault \n")
	}
	context.JSON(http.StatusOK, gin.H{"message": "ok"})
}
