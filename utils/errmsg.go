package utils

var (
	ERROR = 500
	SUCCESS = 200

	// 数据库获取连接错误
	GET_MYSQL_CONNECTION_ERROR = 1001
	// 获取对象数据失败
	GET_MODEL_OBJECT_ERROR = 1002
	// 查无此人
	NOT_FIND_MODEL_ERROR = 1003
	// 用户名已存在
	GET_USERNAME_EXIST_ERROE = 1004
	// 密码有误
	PASSWORD_NOT_RIGHT_ERROR = 1005


	// Token不存在
	ERROR_TOKEN_NOT_EXIST = 2001
	// Token类型有误
	ERROR_TOKEN_TYPE_WRONG = 2002
	// Token错误
	ERROR_TOKEN_WRONG = 2003
	// Token过期
	ERROR_TOKEN_RUNTIME = 2004

	// 验证码绑定错误
	ERROR_BIND_CAPTCHA = 3001
	// 验证码生成错误
	ERROR_CREATE_CAPTCHA = 3002
	// 验证码不匹配
	ERROR_MATCH_CAPTCHA = 3003
	// 验证码无效
	ERROR_CAPTCHA_UNEFFIECT = 3004
)

var errmsg =  map[int]string {
	ERROR: "500错误",
	SUCCESS: "200成功",

	GET_MYSQL_CONNECTION_ERROR: "数据库获取连接错误",
	GET_MODEL_OBJECT_ERROR: "获取数据错误",
	NOT_FIND_MODEL_ERROR: "该用户不存在",
	GET_USERNAME_EXIST_ERROE: "昵称已存在",

	ERROR_TOKEN_NOT_EXIST: "Token不存在",
	ERROR_TOKEN_TYPE_WRONG: "Token类型有误",
	ERROR_TOKEN_WRONG: "Token错误",
	ERROR_TOKEN_RUNTIME: "Token已过期",
	PASSWORD_NOT_RIGHT_ERROR: "密码错误",

	ERROR_BIND_CAPTCHA: "验证码绑定错误",
	ERROR_CREATE_CAPTCHA: "验证码生成错误",
	ERROR_MATCH_CAPTCHA: "验证码不匹配",
	ERROR_CAPTCHA_UNEFFIECT: "此验证码无效",
}

func GetErrMsg(code int) string {
	return errmsg[code]
}