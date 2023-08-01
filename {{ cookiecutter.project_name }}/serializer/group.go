package serializer

import (
	"time"
	"{{ cookiecutter.project_name }}/biz"
)

type GroupSerializer struct {
	Id          int                     `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
	Users       []*UserSerializer       `json:"users"`
	Permissions []*PermissionSerializer `json:"permissions"`
}

func NewGroupSerializer(group *biz.Group) *GroupSerializer {
	gs := &GroupSerializer{
		Id:          group.Id,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}

	if group.IsLoadUsers {
		gs.Users = make([]*UserSerializer, len(group.Users))
		for i, user := range group.Users {
			gs.Users[i] = NewUserSerializer(user)
		}
	}
	if group.IsLoadPermissions {
		gs.Permissions = NewPermissionsSerializer(group.Permissions)
	}
	return gs
}
