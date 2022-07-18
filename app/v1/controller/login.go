package controller

import (
	"github.com/gin-gonic/gin"
	"goserver/app/v1/service"
)

type LoginRouter struct {}
var loginService *service.LoginService

func (ur *LoginRouter) LoginRouterInit(router *gin.RouterGroup) {
	{
		router.POST("/login", loginService.Login)
	}
}