package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"strings"
	"{{ cookiecutter.project_name }}/setting"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			noAuth(c)
			return
		}

		aInfo := strings.Split(auth, " ")
		if len(aInfo) != 2 || aInfo[0] != "Bearer" {
			noAuth(c)
			return
		}
		jwtToken := aInfo[1]

		var (
			tokenInfo *jwt.Token
			err       error
		)

		tokenInfo, err = jwt.ParseWithClaims(jwtToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(setting.AppSetting.JwtSecret), nil
		})

		if err != nil {
			noAuth(c)
			return
		}
		if !tokenInfo.Valid {
			noAuth(c)
			return
		}

		if claims, ok := tokenInfo.Claims.(*jwt.RegisteredClaims); ok && tokenInfo.Valid {
			s, err := strconv.Atoi(claims.Subject)
			if err != nil {
				noAuth(c)
				return
			}
			c.Set("userId", s)
		} else {
			noAuth(c)
			return
		}
		c.Next()
	}
}

func noAuth(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success":      false,
		"errorCode":    401,
		"errorMessage": "请先登录",
		"showType":     2,
	})
	c.Abort()
	return
}
