package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
	"xxvote_server/app/models"
	"xxvote_server/app/tools"
)

type loginUserInfo struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`

	//CaptchaId    string `json:"captcha_id" form:"captcha_id"`
	//CaptchaValue string `json:"captcha_value" form:"captcha_value"`
}

type CreateUserInfo struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

func CreateUser(context *gin.Context) {
	var cUser CreateUserInfo
	if err := context.ShouldBind(&cUser); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    1111,
			Message: err.Error(),
		})
		return
	}
	// 验证创建的用户信息是否有效
	// 所有信息均 不可为空
	if cUser.Name == "" || cUser.Password == "" || cUser.Password2 == "" {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    1111,
			Message: "Information cannot be empty",
			Data:    cUser,
		})
		return
	}
	// 两次输入的密码要 相同
	if cUser.Password != cUser.Password2 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    1111,
			Message: "The passwords entered twice are inconsistent.",
		})
		return
	}
	// 账号名 和 密码要合法 --> 长度 要求[3 ,8] || 不能为纯数字
	if regexp.MustCompile(`^[0-9]+$`).MatchString(cUser.Password) {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    1111,
			Message: "Passwords cannot be pure numbers",
		})
		return
	}

	// 查数据的 检测
	// 用户名 唯一
	if user := models.GetUserByName(cUser.Name); user.Id > 0 {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    1111,
			Message: "user already exists",
		})
		return
	}

	// 都满足 就创建用户
	//验证成功后写入 user表中
	newUser := models.User{
		Uid:        111,
		Name:       cUser.Name,
		Password:   cUser.Password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := models.CreateUser(&newUser); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10010,
			Message: "User creation failed", //这里有风险
		})
		return
	}
	context.JSON(http.StatusOK, tools.OK)
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
	var user *models.User
	if user = models.GetUserByName(userInfo.UserName); user.Id <= 0 || user.Password != userInfo.Password {
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
	if err := tools.FlushSession(context); err != nil {
		fmt.Printf("[DoLogOut] err : %s\n", err.Error())
	}
	// 重定向到login界面
	//context.HTML(http.StatusOK,"login.html",gin.H{})
	context.Redirect(http.StatusFound, "/login")
}
