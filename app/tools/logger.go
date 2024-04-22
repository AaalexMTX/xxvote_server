package tools

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger *logrus.Logger

// NewLogger 创建一个 确定（log类型+写入位置）的log对象
func NewLogger() {
	Logger = logrus.New()
	//设置Logger属性
	//设置报错level
	Logger.SetLevel(logrus.InfoLevel)

	//设置输出位置
	writePos1 := os.Stdout
	writePos2, _ := os.OpenFile("./runtime.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	Logger.SetOutput(io.MultiWriter(writePos1, writePos2))

	//设置上下文
	//设置hook函数
	//设置携带的信息
}
