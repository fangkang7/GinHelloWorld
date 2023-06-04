package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//engine.Use(RequestInfo())
	engine.POST("/kaka", RequestInfo(), func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})
	engine.POST("/kaka2", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	if err := engine.Run(":8089"); err != nil {
		fmt.Println("服务器醋五")
	}
}

func RequestInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求path：", path)
		fmt.Println("请求方式：", method)
	}
}
