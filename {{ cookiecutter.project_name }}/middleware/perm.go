package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{ cookiecutter.project_name }}/service"
)

type PermTool struct {
	us *service.UserService
}

func NewPermTool(us *service.UserService) *PermTool {
	return &PermTool{us: us}
}

func (pt *PermTool) Check(permName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetInt("userId")
		if userId == 0 {
			noPerm(c)
			return
		}
		has, err := pt.us.HasPerm(c, userId, permName)
		if err != nil || !has {
			noPerm(c)
			return
		}
		c.Next()
	}
}

func noPerm(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"success":      false,
		"errorCode":    403,
		"errorMessage": "没有权限",
		"showType":     2,
	})
	c.Abort()
	return
}
