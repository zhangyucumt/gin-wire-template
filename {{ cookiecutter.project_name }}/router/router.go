package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/middleware"
	"{{ cookiecutter.project_name }}/router/api"
)

func NewRouter(api *api.Api, log *logger.Logger, permTool *middleware.PermTool) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Logger(api.GetOperateLogService(), log))
	r.POST("/api/login", api.Login, middleware.SetRequestNameHandler("登录"))
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.Jwt())
	//apiV1.Use(middleware.HasPerm(api.GetUserService()))
	{
		apiV1.POST("/user", permTool.Check("添加用户"), api.CreateUser, middleware.SetRequestNameHandler("创建用户"))
		apiV1.GET("/user", permTool.Check("查看用户"), api.ListUsers, middleware.SetRequestNameHandler("获取用户列表"))
		apiV1.GET("/user/simple", api.SimpleListUsers, middleware.SetRequestNameHandler("获取用户简单信息列表"))
		apiV1.GET("/user/info", api.UserInfo, middleware.SetRequestNameHandler("获取我的信息"))
		apiV1.PUT("/user/reset-my-password", api.ResetMyPassword, middleware.SetRequestNameHandler("重置我的密码"))
		apiV1.PUT("/user/:id", permTool.Check("编辑用户"), api.UpdateUser, middleware.SetRequestNameHandler("更新用户信息"))
		apiV1.DELETE("/user/:id", permTool.Check("删除用户"), api.DeleteUser, middleware.SetRequestNameHandler("删除用户"))
		apiV1.GET("/user/:id", permTool.Check("查看用户"), api.GetUser, middleware.SetRequestNameHandler("获取用户详情"))

		apiV1.GET("/group", permTool.Check("查看用户组"), api.ListGroup, middleware.SetRequestNameHandler("获取组列表"))
		apiV1.POST("/group", permTool.Check("添加用户组"), api.CreateGroup, middleware.SetRequestNameHandler("创建组"))
		apiV1.GET("/group/simple", api.SimpleListGroups, middleware.SetRequestNameHandler("获取组简单信息列表"))
		apiV1.GET("/group/:id", permTool.Check("查看用户组"), api.GetGroup, middleware.SetRequestNameHandler("获取组详情"))
		apiV1.PUT("/group/:id", permTool.Check("编辑用户组"), api.UpdateGroup, middleware.SetRequestNameHandler("更新组信息"))
		apiV1.DELETE("/group/:id", permTool.Check("删除用户组"), api.DeleteGroup, middleware.SetRequestNameHandler("删除组"))

		apiV1.GET("/operate-log", permTool.Check("查看日志"), api.ListOperateLog, middleware.SetRequestNameHandler("获取操作日志列表"))

		apiV1.GET("/permission", api.ListPermission, middleware.SetRequestNameHandler("获取权限列表"))
	}
	return r
}

var ProviderSet = wire.NewSet(NewRouter)
