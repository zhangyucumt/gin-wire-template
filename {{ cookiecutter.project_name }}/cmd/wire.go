//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/data"
	"{{ cookiecutter.project_name }}/middleware"
	"{{ cookiecutter.project_name }}/router"
	"{{ cookiecutter.project_name }}/router/api"
	"{{ cookiecutter.project_name }}/service"
	"{{ cookiecutter.project_name }}/task"
)

func wireApp(log *logrus.Logger) (*app, func(), error) {
	panic(wire.Build(router.ProviderSet, newHttpServer, newApp, data.ProviderSet, service.ProvideSet, biz.ProvideSet, api.ProviderSet, middleware.ProviderSet, task.ProviderSet))
}
