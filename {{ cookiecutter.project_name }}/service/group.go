package service

import (
	"context"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/parser"
	"{{ cookiecutter.project_name }}/serializer"
)

type GroupService struct {
	uc  *biz.GroupUseCase
	log *logger.Logger
}

func NewGroupService(uc *biz.GroupUseCase, log *logger.Logger) *GroupService {
	return &GroupService{
		uc:  uc,
		log: log,
	}
}

func (g *GroupService) Create(ctx *gin.Context, form *parser.CreateGroupParser) error {
	group := &biz.Group{
		Name:        form.Name,
		Description: form.Description,
	}

	if len(form.UserIds) > 0 {
		us := make([]*biz.User, len(form.UserIds))
		for i, id := range form.UserIds {
			us[i] = &biz.User{Id: id}
		}
		group.Users = us
	}

	if len(form.PermissionIds) > 0 {
		ps := make([]*biz.Permission, len(form.PermissionIds))
		for i, id := range form.PermissionIds {
			ps[i] = &biz.Permission{Id: id}
		}
		group.Permissions = ps
	}

	return g.uc.Create(ctx, group)
}

func (g *GroupService) Update(ctx *gin.Context, form *parser.UpdateGroupParser) error {
	group := &biz.Group{
		Id:          form.Id,
		Name:        form.Name,
		Description: form.Description,
	}
	if len(form.UserIds) > 0 {
		us := make([]*biz.User, len(form.UserIds))
		for i, id := range form.UserIds {
			us[i] = &biz.User{Id: id}
		}
		group.Users = us
	}
	if len(form.PermissionIds) > 0 {
		ps := make([]*biz.Permission, len(form.PermissionIds))
		for i, id := range form.PermissionIds {
			ps[i] = &biz.Permission{Id: id}
		}
		group.Permissions = ps
	}
	return g.uc.Update(ctx, group)
}

func (g *GroupService) Delete(ctx *gin.Context, form *parser.DeleteGroupParser) error {
	return g.uc.Delete(ctx, form.Id)
}

func (g *GroupService) Get(ctx *gin.Context, form *parser.GetGroupParser) (*serializer.GroupSerializer, error) {
	group, err := g.uc.Get(ctx, form.Id)
	if err != nil {
		return &serializer.GroupSerializer{}, err
	}
	return serializer.NewGroupSerializer(group), nil
}

func (g *GroupService) List(ctx context.Context) ([]*serializer.GroupSerializer, error) {
	groups, err := g.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	gs := make([]*serializer.GroupSerializer, len(groups))
	for i, group := range groups {
		gs[i] = serializer.NewGroupSerializer(group)
	}
	return gs, nil
}

func (g *GroupService) SimpleList(ctx *gin.Context) ([]*serializer.SimpleSerializer, error) {
	gs, err := g.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	return serializer.NewSimpleListSerializer(gs, "Name", "Id"), nil
}
