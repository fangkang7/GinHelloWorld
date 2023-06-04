package test

import (
	"GinHelloWorld/testProject/tool"
	"fmt"
)

type Common struct {
}

func (c *Common) BaseConfig() {
	_, err := tool.ParseConfig("../config/app.json")
	if err != nil {
		fmt.Println(err.Error())
	}
}
