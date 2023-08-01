package serializer

import (
	"time"
	"{{ cookiecutter.project_name }}/biz"
)

type UserSerializer struct {
	Id          int                     `json:"id"`
	Username    string                  `json:"username"`
	Name        string                  `json:"name"`
	Email       string                  `json:"email"`
	Phone       string                  `json:"phone"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
	LastLogin   *time.Time              `json:"last_login"`
	Groups      []*GroupSerializer      `json:"groups"`
	Permissions []*PermissionSerializer `json:"permissions"`
	IsActive    bool                    `json:"is_active"`
	IsSuperuser bool                    `json:"is_superuser"`
}

func NewUserSerializer(user *biz.User) *UserSerializer {
	us := &UserSerializer{
		Id:          user.Id,
		Username:    user.Username,
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		LastLogin:   user.LastLogin,
		IsActive:    user.IsActive,
		IsSuperuser: user.IsSuperuser,
	}
	if user.IsLoadGroups {
		us.Groups = make([]*GroupSerializer, len(user.Groups))
		for i, group := range user.Groups {
			us.Groups[i] = NewGroupSerializer(group)
		}
	}
	if user.IsLoadPermissions {
		us.Permissions = NewPermissionsSerializer(user.Permissions)
	}
	return us
}

type LoginSerializer struct {
	UserSerializer
	AllPermissions []string `json:"all_permissions"`
	Token          string   `json:"token"`
}

func NewLoginSerializer(user *biz.User, allPermissions []string, token string) *LoginSerializer {
	return &LoginSerializer{
		UserSerializer: *NewUserSerializer(user),
		AllPermissions: allPermissions,
		Token:          token,
	}
}
