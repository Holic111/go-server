package dao

import (
	"goserver/app/v1/model"
	"goserver/utils"
)

type UserDao struct {}

// 获取用户数量
func (ud *UserDao) GetUserCount() int {
	var cnt int
	utils.MysqlDB.Debug().Table("user").Where("deleted_at is null").Count(&cnt)
	return cnt
}

// 分页查询
func (ud *UserDao) GetUsersByPages(pageIndex int, pageSize int) (int, []model.User) {
	var users []model.User

	err := utils.MysqlDB.Debug().Table("user").Limit(pageSize).Offset(pageIndex).Find(&users).Error
	if err != nil {
		return utils.ERROR, nil
	}

	return utils.SUCCESS, users
}

func (ud *UserDao) GetAllUsers() (int, []model.User) {

	var users []model.User

	utils.MysqlDB.Debug().Find(&users)

	return utils.SUCCESS, users
}

// 检查是否有该用户名 true代表有该用户名, false代表没有该用户名
func (ud *UserDao) CheckUser(u model.User) bool {
	var res int
	var user model.User
	username := u.Username

	utils.MysqlDB.Debug().Where("username = ?", username).First(&user).Count(&res)

	return res != 0
}

func (ud *UserDao) AddUser(u model.User) int {
	check := ud.CheckUser(u)
	if check {
		//该用户名已存在
		return utils.GET_USERNAME_EXIST_ERROE
	}

	u.Password = utils.ScryptPw(u.Password)

	var res int
	utils.MysqlDB.Debug().Create(&u).Count(&res)

	if res > 0 {
		return utils.SUCCESS
	}
	return utils.ERROR
}

func (ud *UserDao) GetUserById(id int) (int, model.User) {
	var user model.User
	var count int

	utils.MysqlDB.Debug().Where("id = ?", id).First(&user).Count(&count)

	if count == 0 {
		return utils.NOT_FIND_MODEL_ERROR, user
	}
	return utils.SUCCESS, user
}

func (ud *UserDao) DeleteUser(id int) int {
	code, user := ud.GetUserById(id)

	if code == utils.NOT_FIND_MODEL_ERROR {
		return code
	}

	utils.MysqlDB.Debug().Delete(&user)
	return utils.SUCCESS
}

func (ud *UserDao) UpdateUser(user model.User) int {
	var code int

	//根据username查找
	//check := CheckUser(user)
	//if !check {
	//	code = utils.NOT_FIND_MODEL_ERROR
	//	return code
	//}

	code, u := ud.GetUserById(int(user.ID))
	if code == utils.NOT_FIND_MODEL_ERROR {return code}

	user.Password = utils.ScryptPw(user.Password)

	utils.MysqlDB.Debug().Model(&u).Updates(&user)
	return utils.SUCCESS
}

func (ud *UserDao) CheckPassword(username string, password string) (int, bool) {
	var user model.User
	var count int

	utils.MysqlDB.Debug().Where("username = ? and password = ?", username, password).First(&user).Count(&count)

	if count > 0 {
		return utils.SUCCESS, true
	}
	return utils.PASSWORD_NOT_RIGHT_ERROR, false
}