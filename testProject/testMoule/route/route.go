package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	userGroup := engine.Group("/user")

	userGroup.POST("/add", func(context *gin.Context) {
		context.Writer.Write([]byte("Post请求"))
	})

	userGroup.POST("/update", func(context *gin.Context) {
		context.Writer.WriteString("更新操作")
	})

	engine.Run(":8089")
}
