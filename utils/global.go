package utils

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	MysqlDB *gorm.DB
	Port    string
	Driver        string
	MysqlUser     string
	MysqlPassword string
	IP            string
	DBName string


	Host string
	AppMode string
	PageSize int

	JwtSecret string
	Address string


	RedisDB       int
	RedisPoolSize int
	RedisPassword string
	Redis *redis.Client


	AccessKey string
	SecretKey string
	Bucket string
	QiNiuServer string
)