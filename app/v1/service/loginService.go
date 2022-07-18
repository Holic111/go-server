package service

import (
	"github.com/gin-gonic/gin"
	"goserver/app/v1/model"
	"goserver/middleware"
	"goserver/utils"
)

type LoginService struct {}

func (ls *LoginService) Login(c *gin.Context) {
	var user model.User
	var code int

	username := c.PostForm("username")
	password := c.PostForm("password")

	user.Username = username
	user.Password = password

	//判断验证码
	//captchaId := c.PostForm("captcha_id")
	captcha := c.PostForm("captcha")
	number := c.PostForm("number")
	code = CaptchaVerify(captcha, number)
	if code != utils.SUCCESS {
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"msg": utils.GetErrMsg(code),
		})
		return
	}

	check := userDao.CheckUser(user)
	if check {
		//用户名存在
		var mark bool
		pwd := utils.ScryptPw(user.Password)
		code, mark = userDao.CheckPassword(user.Username, pwd)
		if mark {
			//信息正确，生成Token
			var token string
			token, code = middleware.SetToken(user.Username)
			c.JSON(utils.SUCCESS, gin.H{
				"statusCode": code,
				"token": token,
				"message": utils.GetErrMsg(code),
			})
			return
		}
	}else {
		code = utils.NOT_FIND_MODEL_ERROR
	}

	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"message": utils.GetErrMsg(code),
	})
}