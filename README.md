# xxvote_server
golang vote system

Temporary
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
```

将mysql的建表语句转为 GORM中的结构体
[mysql create table --> GROM struct](https://old.printlove.cn/tools/sql2gorm)
