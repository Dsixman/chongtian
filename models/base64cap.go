package models
import (
	"github.com/mojocn/base64Captcha"
)
func CreateCaptcha() (string,string) {
    //config struct for digits
    //数字验证码配置
    var configD = base64Captcha.ConfigDigit{
        Height:     80,
        Width:      240,
        MaxSkew:    0.7,
        DotCount:   80,
        CaptchaLen: 5,
    }
    idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
    base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
    return idKeyD,base64stringD
}
