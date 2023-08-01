package serializer

import (
	"{{ cookiecutter.project_name }}/biz"
)

type PermissionSerializer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Section string `json:"section"`
}

func NewPermissionSerializer(permission *biz.Permission) *PermissionSerializer {
	return &PermissionSerializer{
		Id:      permission.Id,
		Name:    permission.Name,
		Section: permission.Section,
	}
}

func NewPermissionsSerializer(permissions []*biz.Permission) []*PermissionSerializer {
	ret := make([]*PermissionSerializer, 0, len(permissions))
	for _, permission := range permissions {
		ret = append(ret, NewPermissionSerializer(permission))
	}
	return ret
}
