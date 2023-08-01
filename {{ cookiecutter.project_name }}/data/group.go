package data

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/biz"
	model "{{ cookiecutter.project_name }}/data/ent/group"
)

type groupData struct {
	log  *logger.Logger
	data *Data
}

func NewGroupData(log *logger.Logger, data *Data) biz.GroupRepo {
	return &groupData{log: log, data: data}
}

func (g *groupData) All(ctx context.Context) ([]*biz.Group, error) {
	groups, err := g.data.db.Group.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]*biz.Group, len(groups))
	for i, group := range groups {
		ret[i] = biz.NewGroup(group)
	}
	return ret, nil
}

func (g *groupData) Create(ctx context.Context, group *biz.Group) error {
	gc := g.data.db.Group.Create().SetName(group.Name).SetDescription(group.Description)
	if len(group.Users) > 0 {
		uIds := make([]int, len(group.Users))
		for i, u := range group.Users {
			uIds[i] = u.Id
		}
		gc.AddUserIDs(uIds...)
	}
	if len(group.Permissions) > 0 {
		pIds := make([]int, len(group.Permissions))
		for i, p := range group.Permissions {
			pIds[i] = p.Id
		}
		gc.AddPermissionIDs(pIds...)
	}
	_, err := gc.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (g *groupData) Update(ctx context.Context, group *biz.Group) error {
	gc := g.data.db.Group.UpdateOneID(group.Id).SetName(group.Name).SetDescription(group.Description)
	gc.ClearUsers()
	if len(group.Users) > 0 {
		uIds := make([]int, len(group.Users))
		for i, u := range group.Users {
			uIds[i] = u.Id
		}
		gc.AddUserIDs(uIds...)
	}
	gc.ClearPermissions()
	if len(group.Permissions) > 0 {
		pIds := make([]int, len(group.Permissions))
		for i, p := range group.Permissions {
			pIds[i] = p.Id
		}
		gc.AddPermissionIDs(pIds...)
	}
	_, err := gc.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (g *groupData) Delete(ctx context.Context, id int) error {
	return g.data.db.Group.DeleteOneID(id).Exec(ctx)
}

func (g *groupData) Get(ctx context.Context, id int) (*biz.Group, error) {
	group, err := g.data.db.Group.Query().WithUsers().WithPermissions().Where(model.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return biz.NewGroup(group), nil
}
