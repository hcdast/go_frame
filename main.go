package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"go_frame/handlers"
	"go_frame/log"
	_ "go_frame/models"
	_ "go_frame/task"
)

var (
	globalSessions *session.Manager
)

// 初始化 先main函数执行
func init() {
	// 加载环境变量
	err := godotenv.Load("./conf/app.conf")
	if err != nil {
		log.Logger.Error(err)
	}

	// 设置session TODO
}

// 主程序
func main() {
	beego.ErrorController(&handlers.ErrorController{})
	beego.Run()
}

