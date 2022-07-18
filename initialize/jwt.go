package initialize

import (
	"github.com/spf13/viper"
	"goserver/utils"
)

func JwtInit() {
	v := viper.New()

	v.AddConfigPath("./config/")
	v.SetConfigName("conf")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		return
	}

	utils.JwtSecret = v.GetString("jwt.KeySecret")
}