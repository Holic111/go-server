package service

import (
	"github.com/gin-gonic/gin"
	"goserver/app/v1/model"
	"goserver/utils"
)

type RegisterService struct {}

func (rs *RegisterService) Register(c *gin.Context) {
	var user model.User
	var code int

	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = utils.ERROR
		c.JSON(utils.SUCCESS, gin.H{
			"code": code,
			"message": utils.GetErrMsg(code),
		})
		return
	}

	msg, code := utils.Validate(&user)
	if code != utils.SUCCESS {
		code = utils.ERROR
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"message": msg,
		})
		return
	}

	check := userDao.CheckUser(user)
	if check {
		code = utils.GET_USERNAME_EXIST_ERROE
		c.JSON(utils.SUCCESS, gin.H{
			"code": code,
			"message": utils.GetErrMsg(code),
		})
		return
	}

	code = userDao.AddUser(user)
	c.JSON(utils.SUCCESS, gin.H{
		"code": code,
		"message": utils.GetErrMsg(code),
	})
}