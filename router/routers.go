package router

import c "goserver/app/v1/controller"

type GroupRouters struct {
	c.UserRouters
	c.LoginRouter
	c.RegisterRouter
	c.CaptchaRouter
	c.UploadRouter
}