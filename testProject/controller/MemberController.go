package controller

import (
	"GinHelloWorld/testProject/param"
	"GinHelloWorld/testProject/service"
	"GinHelloWorld/testProject/tool"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Route(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
	// 参数
	/**
	{
	    "phone": "123456",
	    "code": "304955"
	}
	*/
	engine.POST("/api/login", mc.login)
	engine.GET("/api/captcha", mc.captcha)
}

// 生成验证码
func (mc *MemberController) captcha(ctx *gin.Context) {
	captcha, s, err := tool.MakeCaptcha()
	if err != nil {
		// 处理错误
		tool.Failed(ctx, "验证码生成失败")
		return
	}
	captchaResult := tool.CaptchaResult{
		Id:         captcha,
		Base64Blob: s,
	}
	tool.Success(ctx, map[string]interface{}{
		"captcha_result": captchaResult,
	})
}

// 发送短信二维码
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")

	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
	}

	//s := new(service.MemberService)
	//isSend := s.SendCode(phone)

	isSend := (&service.MemberService{}).SendCode(phone)
	if !isSend {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "发送失败",
		})
	}

	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "发送成功",
	})

	tool.Success(context, nil)

	return
}

// 用户登录
func (mc *MemberController) login(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam

	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(context, "参数解析错误")
		return
	}

	us := &service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		tool.Success(context, member)
		return
	}

	tool.Failed(context, "登录失败")
}
