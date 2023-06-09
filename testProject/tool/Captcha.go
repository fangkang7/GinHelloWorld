package tool

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 生成验证码
func MakeCaptcha() *CaptchaResult {
	// 设置自带的 store（可以配置成redis）
	var store = base64Captcha.DefaultMemStore
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}

	//ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()
	//创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	//Generate 生成随机 id、base64 图像字符串
	generate, s, err := captcha.Generate()

	if err != nil {
		return nil
	}

	result := &CaptchaResult{
		Id:         generate,
		Base64Blob: s,
	}

	return result
}
