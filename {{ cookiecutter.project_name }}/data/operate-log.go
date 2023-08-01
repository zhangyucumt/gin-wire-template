package data

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/data/ent"
	model "{{ cookiecutter.project_name }}/data/ent/operatelog"
	"{{ cookiecutter.project_name }}/data/ent/predicate"
)

type operateLogData struct {
	log  *logger.Logger
	data *Data
}

func (o *operateLogData) List(ctx context.Context, page, pageSize int, ordering []ent.OrderFunc, filters ...predicate.OperateLog) ([]*biz.OperateLog, error) {
	page, pageSize = o.data.parserPage(page, pageSize)
	ops, err := o.data.db.OperateLog.Query().WithUser().Where(filters...).Order(ordering...).Offset((page - 1) * pageSize).Limit(pageSize).All(ctx)
	if err != nil {
		return nil, err
	}
	operateLogs := make([]*biz.OperateLog, len(ops))
	for i, op := range ops {
		operateLogs[i] = biz.NewOperateLog(op)
	}
	return operateLogs, nil
}

func (o *operateLogData) Count(ctx context.Context, filters ...predicate.OperateLog) (int, error) {
	return o.data.db.OperateLog.Query().Where(filters...).Count(ctx)
}

func (o *operateLogData) Get(ctx context.Context, id int) (*biz.OperateLog, error) {
	operateLog, err := o.data.db.OperateLog.Query().Where(model.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return biz.NewOperateLog(operateLog), nil
}

func (o *operateLogData) Create(ctx context.Context, operateLog *biz.OperateLog) error {
	_, err := o.data.db.OperateLog.Create().SetName(operateLog.Name).SetPath(operateLog.Path).SetStatusCode(operateLog.StatusCode).SetIP(operateLog.IP).SetMethod(operateLog.Method).SetUserID(operateLog.User.Id).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewOperateLogData(log *logger.Logger, data *Data) biz.OperateLogRepo {
	return &operateLogData{log: log, data: data}
}
