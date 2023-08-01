package biz

import (
	"context"
	"time"
	"{{ cookiecutter.project_name }}/data/ent"
	"{{ cookiecutter.project_name }}/data/ent/predicate"
)

type OperateLog struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Method     string    `json:"method"`
	StatusCode int       `json:"status_code"`
	IP         string    `json:"ip"`
	User       *User     `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewOperateLog(operateLog *ent.OperateLog) *OperateLog {
	if operateLog == nil {
		return nil
	}
	return &OperateLog{
		Id:         operateLog.ID,
		Name:       operateLog.Name,
		Path:       operateLog.Path,
		Method:     operateLog.Method,
		StatusCode: operateLog.StatusCode,
		IP:         operateLog.IP,
		CreatedAt:  operateLog.CreatedAt,
		UpdatedAt:  operateLog.UpdatedAt,
		User:       NewUser(operateLog.Edges.User),
	}
}

type OperateLogRepo interface {
	List(ctx context.Context, page, pageSize int, ordering []ent.OrderFunc, filters ...predicate.OperateLog) ([]*OperateLog, error)
	Count(ctx context.Context, filters ...predicate.OperateLog) (int, error)
	Get(ctx context.Context, id int) (*OperateLog, error)
	Create(ctx context.Context, operateLog *OperateLog) error
}

type OperateLogUseCase struct {
	repo OperateLogRepo
}

func NewOperateLogUseCase(repo OperateLogRepo) *OperateLogUseCase {
	return &OperateLogUseCase{repo: repo}
}

func (uc *OperateLogUseCase) Pagination(ctx context.Context, page, pageSize int, ordering []ent.OrderFunc, filters ...predicate.OperateLog) (int, []*OperateLog, error) {
	count, err := uc.repo.Count(ctx, filters...)
	if err != nil {
		return 0, nil, err
	}

	logs, err := uc.repo.List(ctx, page, pageSize, ordering, filters...)
	if err != nil {
		return 0, nil, err
	}

	return count, logs, nil
}

func (uc *OperateLogUseCase) Get(ctx context.Context, id int) (*OperateLog, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *OperateLogUseCase) Create(ctx context.Context, operateLog *OperateLog) error {
	return uc.repo.Create(ctx, operateLog)
}
