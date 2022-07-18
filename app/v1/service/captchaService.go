package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"goserver/app/v1/model"
	"goserver/utils"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var store = base64Captcha.DefaultMemStore

// 创建验证码
func GenerateCaptchaHandler(c *gin.Context)  {
	var code int
	var param model.ConfigJsonBody

	param.Id = c.PostForm("id")
	param.CaptchaType = c.PostForm("captcha_type")
	param.VerifyValue = c.PostForm("verify_value")
	// 验证码绑定错误
	//if err != nil {
	//	code = utils.ERROR_BIND_CAPTCHA
	//	c.JSON(utils.SUCCESS, gin.H{
	//		"statusCode": code,
	//		"msg": utils.GetErrMsg(code),
	//	})
	//	return
	//}

	var driver base64Captcha.Driver

	// https://captcha.mojotv.cn/ 设置验证码详细参数
	switch param.CaptchaType {
		case "audio" :	//语音
			driver = base64Captcha.DefaultDriverAudio
		case "string":	//字符验证码
			//driver = base64Captcha.NewDriverString()
		case "math":	//数学计算
			//driver = base64Captcha.NewDriverMath()
		case "chinese": //中文验证码
			//driver = base64Captcha.NewDriverChinese()
		default:	//数字验证码
			driver = base64Captcha.DefaultDriverDigit
	}

	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()

	// 生成验证码错误
	if err != nil {
		code = utils.ERROR_CREATE_CAPTCHA
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"msg":  utils.GetErrMsg(code),
		})
		return
	}

	// 将 id 和 number 存入到 redis 中
	number := c.PostForm("number")
	code = SaveCaptcha(id, number)
	if code != utils.SUCCESS {
		fmt.Println("500 ERROR")
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
		})
		return
	}

	code = utils.SUCCESS
	c.JSON(utils.SUCCESS, gin.H{
		"statusCode":      code,
		"data":      b64s,
		"captchaId": id,
		"msg":       utils.GetErrMsg(code),
	})
}

// 校验验证码
func CaptchaVerify(val string, number string) int {
	id := GetCaptchaId(number)
	if id == "" {
		return utils.ERROR_CAPTCHA_UNEFFIECT
	}

	if id == "" || val == "" {
		return utils.ERROR_MATCH_CAPTCHA
	}
	// 验证，同时在内存中清理掉这个图片
	check := store.Verify(id, val, true)
	if check {
		return utils.SUCCESS
	}
	return utils.ERROR_MATCH_CAPTCHA
}

// 根据number从redis中取出验证码信息
func GetCaptchaId(number string) string {
	return utils.Redis.HGet("captcha", number).Val()
}

// 将验证码信息存到redis
func SaveCaptcha(id string, number string) int {
	err := utils.Redis.HSet("captcha", number, id).Err()
	if err != nil {
		return utils.ERROR
	}

	//err = utils.Redis.Expire(number, 60 * time.Second).Err()
	//if err != nil {
	//	return utils.ERROR
	//}

	return utils.SUCCESS
}