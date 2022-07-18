package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"goserver/utils"
)

func QiNiuInit() {
	v := viper.New()

	v.AddConfigPath("./config/")
	v.SetConfigName("conf")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.AccessKey = v.GetString("qiniu.AccessKey")
	utils.SecretKey = v.GetString("qiniu.SecretKey")
	utils.Bucket = v.GetString("qiniu.Bucket")
	utils.QiNiuServer = v.GetString("qiniu.QiNiuServer")

	//fmt.Println(utils.AccessKey)
	//fmt.Println(utils.SecretKey)
	//fmt.Println(utils.Bucket)
	//fmt.Println(utils.QiNiuServer)

	fmt.Println("七牛云初始化完成!")
}