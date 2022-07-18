package controller

import (
	"github.com/gin-gonic/gin"
	"goserver/app/v1/service"
)

type UserRouters struct {}
var userService service.UserService

func (ur *UserRouters) UserRoutersInit(router *gin.RouterGroup) {
	{
		//分页获取用户
		router.GET("/user/page", userService.GetUsersByPage)
		//获取所有用户
		router.GET("/user", userService.GetAllUsers)
		//获取某一个用户
		router.GET("/user/:id", userService.GetUserById)
		//增加用户
		router.POST("/user", userService.AddUser)

		//删除用户
		router.DELETE("/user/:id", userService.DeleteUser)

		//修改用户
		router.PUT("/user", userService.UpdateUser)

		//email
		// router.GET("/user/email", userService.SendEmail)
	}
}