package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"{{ cookiecutter.project_name }}/parser"
)

func (api *Api) Login(c *gin.Context) {
	var form parser.UserLoginParser
	if err := c.ShouldBindJSON(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	ret, err := api.userService.Login(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, ret)
}

func (api *Api) UserInfo(c *gin.Context) {
	//user, exist := c.Get("user")
	//if !exist {
	//	api.error(c, e.RaiseApiException(http.StatusUnauthorized, e.AuthError))
	//	return
	//}
	//api.success(c, user)

	userId := c.GetInt("userId")
	user, err := api.userService.Get(c, &parser.GetUserParser{
		Id: userId,
	})
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, user)

}

func (api *Api) CreateUser(c *gin.Context) {
	var form parser.CreateUserParser
	if err := c.ShouldBindJSON(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	err := api.userService.Create(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}
	api.created(c)
}

func (api *Api) ResetMyPassword(c *gin.Context) {
	var form parser.ResetMyPasswordParser
	if err := c.ShouldBindJSON(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	form.UserId = c.GetInt("userId")
	err := api.userService.ResetMyPassword(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}
	api.updated(c)
}

func (api *Api) UpdateUser(c *gin.Context) {
	uId := c.Param("id")
	// convert cId to int
	uIdInt, err := strconv.Atoi(uId)
	if err != nil {
		api.failure(c, http.StatusBadRequest, "无法识别的参数")
		return
	}

	var form parser.UpdateUserParser
	if err := c.ShouldBindJSON(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	form.Id = uIdInt
	err = api.userService.Update(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}
	api.updated(c)
}

func (api *Api) DeleteUser(c *gin.Context) {
	uId := c.Param("id")
	// convert cId to int
	uIdInt, err := strconv.Atoi(uId)
	if err != nil {
		api.failure(c, http.StatusBadRequest, "无法识别的参数")
		return
	}
	err = api.userService.Delete(c, &parser.DeleteUserParser{
		Id: uIdInt,
	})
	if err != nil {
		api.error(c, err)
		return
	}
	api.updated(c)
}

func (api *Api) GetUser(c *gin.Context) {
	uId := c.Param("id")
	// convert cId to int
	uIdInt, err := strconv.Atoi(uId)
	if err != nil {
		api.failure(c, http.StatusBadRequest, "无法识别的参数")
		return
	}

	user, err := api.userService.Get(c, &parser.GetUserParser{
		Id: uIdInt,
	})
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, user)
}

func (api *Api) ListUsers(c *gin.Context) {
	var form parser.GetUserListParser
	if err := c.ShouldBind(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	users, err := api.userService.List(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, users)
}

func (api *Api) SimpleListUsers(c *gin.Context) {
	users, err := api.userService.SimpleList(c)
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, users)
}
