package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username" validate:"max=20,min=3" label:"昵称"`
	Password string `json:"password" validate:"max=15,min=6" label:"密码"`
	Age int `json:"age" validate:"gte=1,lte=150" label:"年龄"`
	Gender int `json:"gender" validate:"oneof=1 2" label:"性别"`
	RealName string `json:"real_name" validate:"required" label:"真实姓名"`
	PhoneNumber string `json:"phone_number" validate:"len=11" label:"电话号码"`
	Email string `json:"email" validate:"email" label:"电子邮箱"`
	Image string `json:"image" validate:"omitempty" label:"图片"`
	Birth time.Time `json:"birth" validate:"omitempty" label:"出生日期"`
	RoleId int `json:"role_id" validate:"required,gte=1,lte=3" label:"角色id"`
}
