package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"goserver/utils"
)

func RedisInit() {
	v := viper.New()

	v.AddConfigPath("./config/")
	v.SetConfigName("conf")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		return
	}

	utils.Address = v.GetString("redis.Address")
	utils.RedisDB = v.GetInt("redis.RedisDB")
	utils.RedisPoolSize = v.GetInt("redis.RedisPoolSize")
	utils.RedisPassword = v.GetString("redis.RedisPassword")

	rdb := redis.NewClient( &redis.Options {
		Addr: utils.Address,
		Password: utils.RedisPassword,
		DB: utils.RedisDB,
		PoolSize: utils.RedisPoolSize,
	} )

	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败")
		fmt.Println(err)
		return
	}

	utils.Redis = rdb


	fmt.Println("redis连接成功!")
}