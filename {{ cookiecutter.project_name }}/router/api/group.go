package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"{{ cookiecutter.project_name }}/parser"
)

func (api *Api) CreateGroup(c *gin.Context) {
	var form parser.CreateGroupParser
	if err := c.ShouldBindJSON(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	err := api.groupService.Create(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}

	api.created(c)

}

func (api *Api) UpdateGroup(c *gin.Context) {
	gId := c.Param("id")
	// convert gId to int
	gIdInt, err := strconv.Atoi(gId)
	if err != nil {
		api.failure(c, http.StatusBadRequest, "无法识别的参数")
		return
	}
	var form parser.UpdateGroupParser
	if err := c.ShouldBindJSON(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	form.Id = gIdInt
	err = api.groupService.Update(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}
	api.created(c)
}

func (api *Api) DeleteGroup(c *gin.Context) {
	gId := c.Param("id")
	// convert gId to int
	gIdInt, err := strconv.Atoi(gId)
	if err != nil {
		api.failure(c, http.StatusBadRequest, "无法识别的参数")
		return
	}

	err = api.groupService.Delete(c, &parser.DeleteGroupParser{Id: gIdInt})
	if err != nil {
		api.error(c, err)
		return
	}
	api.deleted(c)
}

func (api *Api) GetGroup(c *gin.Context) {
	gId := c.Param("id")
	// convert gId to int
	gIdInt, err := strconv.Atoi(gId)
	if err != nil {
		api.failure(c, http.StatusBadRequest, "无法识别的参数")
		return
	}
	group, err := api.groupService.Get(c, &parser.GetGroupParser{Id: gIdInt})
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, group)
}

func (api *Api) ListGroup(c *gin.Context) {
	groups, err := api.groupService.List(c)
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, groups)
}

func (api *Api) SimpleListGroups(c *gin.Context) {
	users, err := api.groupService.SimpleList(c)
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, users)
}
