package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"goserver/app/v1/model"
	"goserver/utils"
	"time"
)

func DBInit() {

	v := viper.New()

	v.AddConfigPath("./config/")
	v.SetConfigName("conf")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		return
	}

	utils.Port = v.GetString("mysql_database.Port")
	utils.Driver = v.GetString("mysql_database.Driver")
	utils.MysqlUser = v.GetString("mysql_database.MysqlUser")
	utils.MysqlPassword = v.GetString("mysql_database.MysqlPassword")
	utils.IP = v.GetString("mysql_database.IP")
	utils.DBName = v.GetString("mysql_database.DBName")


	db, err := gorm.Open(utils.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.MysqlUser,
		utils.MysqlPassword,
		utils.IP,
		utils.Port,
		utils.DBName,
	))
	if err != nil {
		code := utils.GET_MYSQL_CONNECTION_ERROR
		fmt.Println(utils.GetErrMsg(code))
		return
	}

	//同时最大连接数
	db.DB().SetMaxOpenConns(100)
	//空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)
	//禁用默认表名(+s)
	db.SingularTable(true)
	//自动建表
	db.AutoMigrate(&model.User{}, &model.Role{})
	//连接可复用最大时间
	db.DB().SetConnMaxLifetime(600 * time.Second)

	utils.MysqlDB = db

	if utils.MysqlDB != nil {
		fmt.Println("mysql连接成功!")
	}
	return
}