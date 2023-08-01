package serializer

import (
	"time"
	"{{ cookiecutter.project_name }}/biz"
)

type OperateLogUserSerializer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type OperateLogSerializer struct {
	Id         int                       `json:"id"`
	Name       string                    `json:"name"`
	Path       string                    `json:"path"`
	Method     string                    `json:"method"`
	StatusCode int                       `json:"status_code"`
	IP         string                    `json:"ip"`
	User       *OperateLogUserSerializer `json:"user"`
	CreatedAt  time.Time                 `json:"created_at"`
	UpdatedAt  time.Time                 `json:"updated_at"`
}

func NewOperateLogSerializer(ol *biz.OperateLog) *OperateLogSerializer {
	user := ol.User
	var u OperateLogUserSerializer
	if user != nil {
		u = OperateLogUserSerializer{Id: ol.User.Id, Name: ol.User.Name, Username: ol.User.Username}
	}

	return &OperateLogSerializer{
		Id:         ol.Id,
		Name:       ol.Name,
		Path:       ol.Path,
		Method:     ol.Method,
		StatusCode: ol.StatusCode,
		IP:         ol.IP,
		User:       &u,
		CreatedAt:  ol.CreatedAt,
		UpdatedAt:  ol.UpdatedAt,
	}
}
