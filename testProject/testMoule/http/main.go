package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//engine := gin.Default()
	//
	//engine.GET("/hello", func(context *gin.Context) {
	//	fmt.Println("请求路径:", context.FullPath())
	//	context.Writer.Write([]byte("Hello , gin\n"))
	//})
	//
	//if err := engine.Run(":8090"); err != nil {
	//	log.Fatal(err.Error())
	//}

	engine := gin.Default()

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "kaka")
		fmt.Println(name)

		context.Writer.Write([]byte("hello ," + name))
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		username := context.PostForm("name")
		password := context.PostForm("password")

		param, _ := context.GetPostForm("lala")
		fmt.Println(param)

		if username != "kaka" && password != "123456" {
			context.Writer.Write([]byte("您的账户密码错误，请重试！"))
		}
	})

	engine.GET("/kaka1", func(context *gin.Context) {
		context.Writer.Write([]byte("我是get请求！"))
	})

	engine.POST("/kaka2", func(context *gin.Context) {
		name, exist := context.GetPostForm("kaka")
		if exist {
			fmt.Println(name)
		}
		context.Writer.Write([]byte("我是post请求"))
	})

	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	}
}
