package biz

import (
	"context"
	"time"
	"{{ cookiecutter.project_name }}/data/ent"
)

type Group struct {
	Id          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Users       []*User       `json:"users"`
	Permissions []*Permission `json:"permissions"`

	IsLoadUsers       bool `json:"-"`
	IsLoadPermissions bool `json:"-"`
}

func NewGroup(group *ent.Group) *Group {
	if group == nil {
		return nil
	}
	g := &Group{
		Id:          group.ID,
		Name:        group.Name,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}
	users, ue := group.Edges.UsersOrErr()
	permissions, pe := group.Edges.PermissionsOrErr()

	g.IsLoadPermissions = pe == nil
	g.IsLoadUsers = ue == nil

	g.Permissions = make([]*Permission, len(permissions))
	for i, permission := range permissions {
		g.Permissions[i] = NewPermission(permission)
	}

	g.Users = make([]*User, len(users))
	for i, user := range users {
		g.Users[i] = NewUser(user)
	}
	return g
}

type GroupRepo interface {
	//List(ctx context.Context, page, pageSize int, filters ...predicate.Group) ([]*Group, error)
	All(ctx context.Context) ([]*Group, error)
	//Count(ctx context.Context, filters ...predicate.Group) (int, error)
	Create(ctx context.Context, group *Group) error
	Update(ctx context.Context, group *Group) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*Group, error)
}

type GroupUseCase struct {
	repo GroupRepo
}

func NewGroupUseCase(repo GroupRepo) *GroupUseCase {
	return &GroupUseCase{repo: repo}
}

func (uc *GroupUseCase) List(ctx context.Context) ([]*Group, error) {
	return uc.repo.All(ctx)
}

func (uc *GroupUseCase) Create(ctx context.Context, group *Group) error {
	return uc.repo.Create(ctx, group)
}

func (uc *GroupUseCase) Update(ctx context.Context, group *Group) error {
	return uc.repo.Update(ctx, group)
}

func (uc *GroupUseCase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *GroupUseCase) Get(ctx context.Context, id int) (*Group, error) {
	return uc.repo.Get(ctx, id)
}
