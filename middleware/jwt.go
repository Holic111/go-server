package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goserver/utils"
	"strings"
	"time"
)

//密钥参数
var JwtKey = []byte(utils.JwtSecret)

//接收参数
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成Token
func SetToken(username string) (string, int) {
	//时间有效区间
	expireTime := time.Now().Add(10 * time.Hour)

	setClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "gin-server",
		},
	}

	//返回 *Token，Token是一个结构体
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)

	//将Token指针转为已经签好名的token字符串
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", utils.ERROR
	}
	return token, utils.SUCCESS
}

//验证中间件
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(
		token,
		&MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)

	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, utils.SUCCESS
	}else {
		return nil, utils.ERROR
	}
}

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")

		code := utils.SUCCESS
		if tokenHeader == "" {
			code = utils.ERROR_TOKEN_NOT_EXIST
			c.JSON(utils.SUCCESS, gin.H{
				"statusCode": code,
				"message": utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHeader," ",2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = utils.ERROR_TOKEN_TYPE_WRONG
			c.JSON(utils.SUCCESS, gin.H{
				"code": code,
				"message": utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		key, tCode := CheckToken(checkToken[1])
		if tCode == utils.ERROR {
			code = utils.ERROR_TOKEN_WRONG
			c.JSON(utils.SUCCESS, gin.H{
				"code": code,
				"message": utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > key.ExpiresAt {
			//过期
			code = utils.ERROR_TOKEN_RUNTIME
			c.JSON(utils.SUCCESS, gin.H{
				"code": code,
				"message": utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}