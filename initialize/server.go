package initialize

import (
	"github.com/spf13/viper"
	"goserver/utils"
)

func ServerInit() {
	v := viper.New()

	v.AddConfigPath("./config/")
	v.SetConfigName("conf")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		return 
	}

	utils.Host = v.GetString("server.Host")
	utils.AppMode = v.GetString("server.AppMode")
	utils.PageSize = v.GetInt("server.PageSize")
}