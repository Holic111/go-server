package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	RoleId int `json:"role_id" validate:"lt=1,gt=3" label:"角色id"`
	RoleName string `json:"role_name" validate:"min=2,max=8" label:"角色名称"`
}