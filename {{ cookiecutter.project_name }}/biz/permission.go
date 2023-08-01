package biz

import (
	"context"
	"{{ cookiecutter.project_name }}/data/ent"
)

type Permission struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Section string `json:"section"`
}

func NewPermission(permission *ent.Permission) *Permission {
	if permission == nil {
		return nil
	}
	return &Permission{
		Id:      permission.ID,
		Name:    permission.Name,
		Section: permission.Section,
	}
}

type PermissionRepo interface {
	All(ctx context.Context) ([]*Permission, error)
	Count(ctx context.Context) (int, error)
	CreateOrUpdate(ctx context.Context, permission *Permission) error
}

type PermissionUseCase struct {
	repo PermissionRepo
}

func NewPermissionUseCase(repo PermissionRepo) *PermissionUseCase {
	return &PermissionUseCase{repo: repo}
}

func (uc *PermissionUseCase) List(ctx context.Context) ([]*Permission, error) {
	permissions, err := uc.repo.All(ctx)
	if err != nil {
		return nil, err
	}

	//ret := make(map[string][]*Permission)
	//for _, permission := range permissions {
	//	ele, ok := ret[permission.Section]
	//	if !ok {
	//		ret[permission.Section] = []*Permission{permission}
	//	} else {
	//		ret[permission.Section] = append(ele, permission)
	//	}
	//}
	return permissions, nil
}

func (uc *PermissionUseCase) CreateOrUpdate(ctx context.Context, permission *Permission) error {
	return uc.repo.CreateOrUpdate(ctx, permission)
}
