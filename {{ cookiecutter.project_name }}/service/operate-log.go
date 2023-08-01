package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"strings"
	"time"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/data/ent"
	"{{ cookiecutter.project_name }}/data/ent/operatelog"
	"{{ cookiecutter.project_name }}/data/ent/predicate"
	"{{ cookiecutter.project_name }}/data/ent/user"
	"{{ cookiecutter.project_name }}/parser"
	"{{ cookiecutter.project_name }}/serializer"
)

type OperateLogService struct {
	uc  *biz.OperateLogUseCase
	log *logger.Logger
}

func NewOperateLogService(uc *biz.OperateLogUseCase, log *logger.Logger) *OperateLogService {
	return &OperateLogService{
		uc:  uc,
		log: log,
	}
}

func (o *OperateLogService) List(ctx *gin.Context, form *parser.ListOperateLogParser) (*serializer.PaginationSerializer, error) {

	filters := make([]predicate.OperateLog, 0)
	if form.UserId > 0 {
		filters = append(filters, operatelog.HasUserWith(user.ID(form.UserId)))
	}

	if form.CreatedAtGt != "" {
		parse, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00+08:00", form.CreatedAtGt))
		if err == nil {
			filters = append(filters, operatelog.CreatedAtGT(parse))
		} else {
			o.log.Errorf("parse %v error: %v", form.CreatedAtGt, err)
		}
	}

	if form.CreatedAtLt != "" {
		//filters = append(filters, operatelog.CreatedAtLTE(*form.CreatedAtLt))
		parse, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59+08:00", form.CreatedAtLt))
		if err == nil {
			filters = append(filters, operatelog.CreatedAtLTE(parse))
		} else {
			o.log.Errorf("parse %v error: %v", form.CreatedAtLt, err)
		}
	}

	if form.Path != "" {
		filters = append(filters, operatelog.Path(form.Path))
	}

	if form.Name != "" {
		filters = append(filters, operatelog.Name(form.Name))
	}

	if form.Ip != "" {
		filters = append(filters, operatelog.IP(form.Ip))
	}

	if form.Method != "" {
		filters = append(filters, operatelog.Method(form.Method))
	}

	if form.IsLoginLog == "" || form.IsLoginLog == "false" {
		filters = append(filters, operatelog.PathNEQ("/api/login"))
	}

	if form.IsLoginLog == "true" {
		filters = append(filters, operatelog.Path("/api/login"))
	}

	ords := make([]ent.OrderFunc, 0)
	if form.Ordering != "" {
		ordings := strings.Split(form.Ordering, ",")
		for _, ording := range ordings {
			if strings.HasPrefix(ording, "-") {
				ords = append(ords, ent.Desc(ording[1:]))
			} else {
				ords = append(ords, ent.Asc(ording))
			}
		}
	} else {
		ords = append(ords, ent.Desc("created_at"))
	}

	total, logs, err := o.uc.Pagination(ctx, form.Page, form.PageSize, ords, filters...)
	if err != nil {
		return nil, err
	}
	ols := make([]*serializer.OperateLogSerializer, len(logs))
	for i, log := range logs {
		ols[i] = serializer.NewOperateLogSerializer(log)
	}
	return serializer.NewPaginationSerializer(total, ols), nil
}

func (o *OperateLogService) Create(ctx *gin.Context, form *parser.CreateOperateLogParser) error {
	ol := &biz.OperateLog{
		Name:       form.Name,
		Path:       form.Path,
		Method:     form.Method,
		StatusCode: form.StatusCode,
		IP:         form.IP,
		User:       &biz.User{Id: form.UserId},
	}
	return o.uc.Create(ctx, ol)
}
