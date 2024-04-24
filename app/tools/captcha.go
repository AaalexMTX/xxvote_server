package tools

import "github.com/mojocn/base64Captcha"

type CaptchaData struct {
	CaptchaId string `json:"captcha_id"`
	Data      string `json:"data"`
}

// 数字驱动——驱动设置 验证码的属性
var digitDriver = base64Captcha.DriverDigit{
	Height:   60,  //生成图片高度
	Width:    150, //生成图片宽度
	Length:   4,   //验证码长度
	MaxSkew:  0.5, //文字的倾斜度 越大倾斜越狠，越不容易看懂
	DotCount: 1,   //背景的点数，越大，字体越模糊
}

// 存在内存中 -- 内存驱动，相关数据会存在内存空间里
var storeCaptcha = base64Captcha.DefaultMemStore

func CaptchaGenerate() (CaptchaData, error) {
	var ret CaptchaData
	captcha := base64Captcha.NewCaptcha(&digitDriver, storeCaptcha)
	id, bs64, _, err := captcha.Generate()
	if err != nil {
		return ret, err
	}

	ret.CaptchaId = id
	ret.Data = bs64
	return ret, nil
}

func CaptchaVerify(data CaptchaData) bool {
	return storeCaptcha.Verify(data.CaptchaId, data.Data, true)
}
