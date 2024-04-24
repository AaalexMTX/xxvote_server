# xxvote_server
golang vote system

Import package records
```shell
# GIN
go get "github.com/gin-gonic/gin"
# GORM
go get "github.com/go-sql-driver/mysql" #  MySQL 数据库的 Go 语言驱动程序 _ "" 因为不需要使用，但要引入
go get "gorm.io/driver/mysql"  # GORM 对 MySQL 数据库的专用驱动程序，负责与 MySQL 数据库进行交互
go get "gorm.io/gorm"          # GORM 的核心包，操作数据库
# Viper
go get "github.com/spf13/viper"
# Logurs
go get "github.com/sirupsen/logrus"
# Session
go get -u "github.com/gorilla/sessions"         #cookie
go get -u "github.com/rbcervilla/redisstore/v9" #redis
# captcha / base64Captcha
go get -u "github.com/mojocn/base64Captcha"
```

将mysql的建表语句转为 GORM中的结构体<br>
[mysql create table --> GROM struct](https://old.printlove.cn/tools/sql2gorm)

## 资料
[bili_使用Gin框架搭建项目](https://www.bilibili.com/video/BV1Nj421U7gj/?share_source=copy_web&vd_source=e8c6412141d88fdf27ca603b433b24bf)