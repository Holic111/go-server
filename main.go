package main

import (
	in "goserver/initialize"
)

// @title go-server
// @version 1.0
// @description 测试swagger
// @host localhost:8080
// BasePath /api/v1
func main() {
	in.DBInit()
	in.RedisInit()
	in.QiNiuInit()
	in.ServerInit()
	in.RoutersInit()
	in.JwtInit()
}
