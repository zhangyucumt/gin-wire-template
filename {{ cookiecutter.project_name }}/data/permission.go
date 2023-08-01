package data

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/data/ent"
	"{{ cookiecutter.project_name }}/data/ent/permission"
)

type permissionData struct {
	log  *logger.Logger
	data *Data
}

func NewPermissionData(log *logger.Logger, data *Data) biz.PermissionRepo {
	return &permissionData{log: log, data: data}
}

func (p *permissionData) All(ctx context.Context) ([]*biz.Permission, error) {
	ps, err := p.data.db.Permission.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	permissions := make([]*biz.Permission, len(ps))

	for i, p := range ps {
		permissions[i] = biz.NewPermission(p)
	}
	return permissions, nil
}

func (p *permissionData) Count(ctx context.Context) (int, error) {
	return p.data.db.Permission.Query().Count(ctx)
}

func (p *permissionData) CreateOrUpdate(ctx context.Context, perm *biz.Permission) error {
	tx, err := p.data.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = p.data.rollback(tx, err, recover())
	}()

	o, err := tx.Permission.Query().Where(permission.Name(perm.Name)).ForUpdate().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			_, err = tx.Permission.Create().SetName(perm.Name).SetSection(perm.Section).Save(ctx)
			if err != nil {
				return err
			}
		}
		return err
	}
	_, err = tx.Permission.UpdateOne(o).SetName(perm.Name).SetSection(perm.Section).Save(ctx)
	return err
}
