package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (hello *HelloController) Route(engine *gin.Engine) {
	engine.POST("/hello", hello.Hello)
}

func (hello *HelloController) Hello(context *gin.Context) {
	context.Writer.WriteString("hello world")
}
