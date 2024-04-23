package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// 创建一个类似密钥，保存session的对象
// 1. 存在cookie里
var store = sessions.NewCookieStore([]byte("xxvote_server"))

// 2. 存在redis里，要配置redis
// var store2 *redisstore.RedisStore
// 不同类型的 session 键，可自定义
var sessionName = "session-name"

func GetSession(context *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(context.Request, sessionName)
	//values 不定长
	return session.Values
}

func SetSession(context *gin.Context, name string, id int) error {
	// 给session对象，并设置session字段
	session, _ := store.Get(context.Request, sessionName)
	session.Values["name"] = name
	session.Values["id"] = id
	return session.Save(context.Request, context.Writer)
}

func FlushSession(context *gin.Context) error {
	session, _ := store.Get(context.Request, sessionName)
	session.Values["name"] = ""
	session.Values["id"] = 0
	return session.Save(context.Request, context.Writer)
}
