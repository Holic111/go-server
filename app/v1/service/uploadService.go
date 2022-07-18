package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"goserver/utils"
	"mime/multipart"
	"strconv"
)

type UploadService struct {}

func (u *UploadService) UploadFile(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(utils.SUCCESS, gin.H{
			"msg": "id错误",
		})
		return
	}

	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size

	if file == nil {
		c.JSON(utils.SUCCESS, gin.H{
			"msg": "请先上传文件",
		})
		return
	}
	code, user := userDao.GetUserById(id)
	if code != utils.SUCCESS {
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"msg": utils.GetErrMsg(code),
		})
		return
	}

	imgUrl, code := UploadFileToQiNiu(file, fileSize)
	if code != utils.SUCCESS {
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"msg": utils.GetErrMsg(code),
		})
		return
	}

	user.Image = imgUrl

	code = userDao.UpdateUser(user)
	if code != utils.SUCCESS {
		c.JSON(utils.SUCCESS, gin.H{
			"statusCode": code,
			"msg": utils.GetErrMsg(code),
		})
		return
	}

	c.JSON(utils.SUCCESS, gin.H{
		"statusCode": code,
		"imgUrl": imgUrl,
		"msg": utils.GetErrMsg(code),
	})
}

func UploadFileToQiNiu(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: utils.Bucket,
	}

	mac := qbox.NewMac(utils.AccessKey, utils.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone: &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS: false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)

	if err != nil {
		return "", utils.ERROR
	}

	url := utils.QiNiuServer + "/" + ret.Key
	return url, utils.SUCCESS
}