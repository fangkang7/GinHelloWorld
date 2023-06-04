package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	engine.POST("/hellobyte", func(context *gin.Context) {
		context.Writer.Write([]byte("请求路径为：" + context.FullPath()))
	})

	engine.POST("/helloWrite", func(context *gin.Context) {
		context.Writer.WriteString("请求路径为：" + context.FullPath())
	})

	engine.POST("/helloJson", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    context.FullPath(),
		})
	})

	engine.POST("/helloJsonStruct", func(context *gin.Context) {
		response := Response{
			Code:    1,
			Message: "ok",
			Data:    context.FullPath(),
		}
		context.JSON(200, response)
	})

	engine.Run(":8089")
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
