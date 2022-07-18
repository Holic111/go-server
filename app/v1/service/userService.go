package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"goserver/app/v1/dao"
	"goserver/app/v1/model"
	"goserver/utils"
	"net/smtp"
	"net/textproto"
	"strconv"
)

type UserService struct {}

var userDao dao.UserDao

// 分页获取用户
func (us *UserService) GetUsersByPage(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize := utils.PageSize
	var page int

	if pageNum <= 0 {
		pageNum = 1
	}

	count := userDao.GetUserCount()

	if count % pageSize == 0 {
		page = count / pageSize
	}else {
		page = count / pageSize + 1
	}

	var pageIndex int
	if pageNum > page {
		pageNum = page
	}

	pageIndex = (pageNum - 1) * pageSize

	fmt.Println(pageIndex)
	fmt.Println(pageNum)
	fmt.Println(pageSize)
	fmt.Println(page)

	code, users  := userDao.GetUsersByPages(pageIndex, pageSize)
	if code != utils.SUCCESS {
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"msg": utils.GetErrMsg(code),
		})
		return
	}
	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"msg": utils.GetErrMsg(code),
		"data": users,
		"pageNum": pageNum,
		"pageSize": pageSize,
		"page": page,
		"pageIndex": pageIndex,
		"count": count,
	})
}

// 获取所有用户
func (us *UserService) GetAllUsers(c *gin.Context) {
	code, users := userDao.GetAllUsers()

	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"data": users,
		"message": utils.GetErrMsg(code),
	})
}

func (us *UserService) AddUser(c *gin.Context) {
	var user model.User
	var code int

	err := c.ShouldBindJSON(&user)
	if err != nil {
		code = utils.GET_MODEL_OBJECT_ERROR
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"message": utils.GetErrMsg(code),
		})
		fmt.Println(err)
		return
	}

	// 传入结构体地址
	msg, code := utils.Validate(&user)
	if code != utils.SUCCESS {
		code = utils.ERROR
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"message": msg,
		})
		return
	}

	code = userDao.AddUser(user)

	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"message": utils.GetErrMsg(code),
	})
}

func (us *UserService) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var code int
	if err != nil {
		code = utils.ERROR
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	code, user := userDao.GetUserById(id)
	if code == utils.NOT_FIND_MODEL_ERROR {
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"message": utils.GetErrMsg(code),
		})
		return
	}
	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"data": user,
		"message": utils.GetErrMsg(code),
	})
}

func (us *UserService) DeleteUser(c *gin.Context) {
	var code int
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		code = utils.ERROR
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"message": utils.GetErrMsg(code),
		})
		return
	}

	code = userDao.DeleteUser(id)
	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"message": utils.GetErrMsg(code),
	})
}

func (us *UserService) UpdateUser(c *gin.Context) {
	var user model.User
	var code int

	err := c.ShouldBindJSON(&user)
	fmt.Println(user)
	if err != nil {
		code = utils.ERROR
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"message": utils.GetErrMsg(code),
		})
		return
	}

	code = userDao.UpdateUser(user)
	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"message": utils.GetErrMsg(code),
	})
}

// 发送email
func (us *UserService) SendEmail(c *gin.Context) {
	e := &email.Email {
		To: []string{"2270353958@qq.com"},
		From: "lfz",
		Subject: "=====Subject=====",
		Text: []byte("====Text-LFZEmail===="),
		HTML: []byte("<h3>Html Supported!</h3>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send("smtp.qq.com:587",smtp.PlainAuth("","227035398@qq.com","nhlgvywuhdcgeagc","smtp.qq.com"))
	if err != nil {
		c.JSON(utils.SUCCESS, gin.H{
			"errMsg": err.Error(),
		})
		return 
	}
}
