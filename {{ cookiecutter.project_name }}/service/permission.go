package service

import (
	"context"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/serializer"
)

type PermissionService struct {
	uc  *biz.PermissionUseCase
	log *logger.Logger
}

func NewPermissionService(uc *biz.PermissionUseCase, log *logger.Logger) *PermissionService {
	p := &PermissionService{
		uc:  uc,
		log: log,
	}
	p.initPerms()
	return p
}

func (p *PermissionService) initPerms() {
	ctx := context.Background()
	permissions := []biz.Permission{
		{Name: "查看物理设备监控", Section: "监控指标"},
		{Name: "查看交换机设备监控", Section: "监控指标"},
		{Name: "查看存储平台监控", Section: "监控指标"},
		{Name: "查看云平台监控", Section: "监控指标"},
		{Name: "查看告警规则", Section: "告警管理"},
		{Name: "添加告警规则", Section: "告警管理"},
		{Name: "编辑告警规则", Section: "告警管理"},
		{Name: "删除告警规则", Section: "告警管理"},
		{Name: "查看告警消息", Section: "告警管理"},
		{Name: "查看告警通知配置", Section: "告警管理"},
		{Name: "设置告警通知配置", Section: "告警管理"},
		//{Name: "查看告警通知接收端", Section: "告警管理"},
		//{Name: "添加告警通知接收端", Section: "告警管理"},
		//{Name: "编辑告警通知接收端", Section: "告警管理"},
		//{Name: "删除告警通知接收端", Section: "告警管理"},
		{Name: "获取告警通知", Section: "告警管理"},
		{Name: "查看用户", Section: "用户"},
		{Name: "添加用户", Section: "用户"},
		{Name: "编辑用户", Section: "用户"},
		{Name: "删除用户", Section: "用户"},
		{Name: "查看用户组", Section: "用户组"},
		{Name: "添加用户组", Section: "用户组"},
		{Name: "编辑用户组", Section: "用户组"},
		{Name: "删除用户组", Section: "用户组"},
		{Name: "查看日志", Section: "日志管理"},
		{Name: "查看数据源", Section: "数据源管理"},
		{Name: "添加数据源", Section: "数据源管理"},
		{Name: "编辑数据源", Section: "数据源管理"},
		{Name: "删除数据源", Section: "数据源管理"},
		{Name: "查看交换机", Section: "交换机管理"},
		{Name: "添加交换机", Section: "交换机管理"},
		{Name: "编辑交换机", Section: "交换机管理"},
		{Name: "删除交换机", Section: "交换机管理"},

		{Name: "查看IPMI配置", Section: "IPMI管理"},
		{Name: "添加IPMI配置", Section: "IPMI管理"},
		{Name: "编辑IPMI配置", Section: "IPMI管理"},
		{Name: "删除IPMI配置", Section: "IPMI管理"},
	}

	for _, perm := range permissions {
		err := p.uc.CreateOrUpdate(ctx, &perm)
		if err != nil {
			panic(err)
		}
	}
}

func (p *PermissionService) List(ctx *gin.Context) ([]*serializer.PermissionSerializer, error) {
	ps, err := p.uc.List(ctx)
	if err != nil {
		return nil, err
	}

	ret := make([]*serializer.PermissionSerializer, len(ps))
	for i, perm := range ps {
		ret[i] = serializer.NewPermissionSerializer(perm)
	}
	return ret, nil
}
