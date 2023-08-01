package api

import (
	"github.com/gin-gonic/gin"
)

func (api *Api) ListPermission(c *gin.Context) {
	ps, err := api.permissionService.List(c)
	if err != nil {
		api.error(c, err)
		return
	}
	api.success(c, ps)
}
