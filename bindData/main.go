package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.POST("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var student Student
		err := context.ShouldBindQuery(&student)

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(student.Name)
		fmt.Println(student.Class)
	})
	engine.Run(":8089")
}

type Student struct {
	Name  string `form:"name"`
	Class string `form:"class"`
}
