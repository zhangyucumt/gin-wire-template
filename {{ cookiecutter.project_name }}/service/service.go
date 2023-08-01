package service

import "github.com/google/wire"

var ProvideSet = wire.NewSet(NewUserService, NewPermissionService, NewGroupService, NewOperateLogService)
