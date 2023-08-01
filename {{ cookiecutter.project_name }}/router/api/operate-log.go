package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{ cookiecutter.project_name }}/parser"
)

func (api *Api) ListOperateLog(c *gin.Context) {
	var form parser.ListOperateLogParser
	if err := c.ShouldBind(&form); err != nil {
		api.failure(c, http.StatusBadRequest, err.Error())
		return
	}
	data, err := api.operateLogService.List(c, &form)
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, data)
}
