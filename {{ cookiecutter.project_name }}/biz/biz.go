package biz

import "github.com/google/wire"

var ProvideSet = wire.NewSet(NewUserUseCase, NewGroupUseCase, NewOperateLogUseCase, NewPermissionUseCase)
