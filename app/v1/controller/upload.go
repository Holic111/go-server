package controller

import (
	"github.com/gin-gonic/gin"
	"goserver/app/v1/service"
)

type UploadRouter struct {}
var upload service.UploadService

func (u *UploadRouter) UploadRouterInit(router *gin.RouterGroup) {
	{
		router.POST("/upload", upload.UploadFile)
	}
}