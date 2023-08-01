package biz

import (
	"context"
	"github.com/toolkits/pkg/slice"
	"time"
	"{{ cookiecutter.project_name }}/data/ent"
	"{{ cookiecutter.project_name }}/data/ent/predicate"
)

type User struct {
	Id                int           `json:"id"`
	Username          string        `json:"username,omitempty"`
	Name              string        `json:"name"`
	Password          string        `json:"-"`
	HashedPassword    string        `json:"-"`
	Email             string        `json:"email,omitempty"`
	Phone             string        `json:"phone,omitempty"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	LastLogin         *time.Time    `json:"last_login"`
	Groups            []*Group      `json:"groups"`
	Permissions       []*Permission `json:"permissions"`
	IsLoadGroups      bool          `json:"-"`
	IsLoadPermissions bool          `json:"-"`
	IsActive          bool          `json:"is_active"`
	IsSuperuser       bool          `json:"is_superuser"`
}

func NewUser(user *ent.User) *User {
	if user == nil {
		return nil
	}
	u := &User{
		Id:             user.ID,
		Name:           user.Name,
		Username:       user.Username,
		HashedPassword: user.Password,
		Email:          user.Email,
		Phone:          user.Phone,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		LastLogin:      user.LastLogin,
		IsActive:       user.IsActive,
		IsSuperuser:    user.IsSuperuser,
	}
	groups, gs := user.Edges.GroupsOrErr()
	permissions, ps := user.Edges.PermissionsOrErr()
	u.Groups = make([]*Group, len(groups))
	for i, group := range groups {
		u.Groups[i] = NewGroup(group)
	}
	u.Permissions = make([]*Permission, len(permissions))
	for i, permission := range permissions {
		u.Permissions[i] = NewPermission(permission)
	}
	u.IsLoadGroups = gs == nil
	u.IsLoadPermissions = ps == nil
	return u
}

type UserRepo interface {
	Login(ctx context.Context, username, password string) (*User, error)
	List(ctx context.Context, page, pageSize int, ordering []ent.OrderFunc, filters ...predicate.User) ([]*User, error)
	All(ctx context.Context) ([]*User, error)
	Count(ctx context.Context, filters ...predicate.User) (int, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*User, error)
	GetUserAllPermissions(ctx context.Context, id int) ([]string, error)
	ValidatePassword(hashedPassword, password string) bool
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Login(ctx context.Context, username, password string) (*User, error) {
	return uc.repo.Login(ctx, username, password)
}

func (uc *UserUseCase) Pagination(ctx context.Context, page, pageSize int, ordering []ent.OrderFunc, filters ...predicate.User) (int, []*User, error) {
	count, err := uc.repo.Count(ctx, filters...)
	if err != nil {
		return 0, nil, err
	}

	users, err := uc.repo.List(ctx, page, pageSize, ordering, filters...)
	if err != nil {
		return 0, nil, err
	}
	return count, users, nil
}

func (uc *UserUseCase) All(ctx context.Context) ([]*User, error) {
	return uc.repo.All(ctx)
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) error {
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) error {
	return uc.repo.Update(ctx, user)
}

func (uc *UserUseCase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *UserUseCase) Get(ctx context.Context, id int) (*User, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *UserUseCase) HasPerm(ctx context.Context, id int, permName string) (bool, error) {
	permissions, err := uc.repo.GetUserAllPermissions(ctx, id)
	if err != nil {
		return false, err
	}
	return slice.ContainsString(permissions, permName), nil
}

func (uc *UserUseCase) GetUserAllPermissions(ctx context.Context, id int) ([]string, error) {
	return uc.repo.GetUserAllPermissions(ctx, id)
}

func (uc *UserUseCase) ValidatePassword(hashedPassword, password string) bool {
	return uc.repo.ValidatePassword(hashedPassword, password)
}
