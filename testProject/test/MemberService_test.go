package test

import (
	"GinHelloWorld/testProject/tool"
	"fmt"
	"testing"
)

func TestMemberServiceSendCode(t *testing.T) {
	(&Common{}).BaseConfig()
	//memberService := &service.MemberService{}
	//code := memberService.SendCode("13020733")
	//fmt.Println(code)

	configSms := tool.GetConfig()
	fmt.Println(configSms.Sms.AppSecret)
}
