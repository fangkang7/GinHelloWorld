package main

import (
	"GinHelloWorld/testProject/controller"
	"GinHelloWorld/testProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

/*
*
注册路由
*/
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Route(router)
	new(controller.MemberController).Route(router)
}
