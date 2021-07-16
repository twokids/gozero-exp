package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"zeroService/api/internal/config"
	"zeroService/api/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserinfoModel //新加
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserinfoModel(sqlx.NewMysql(c.DataSource)), //新加
	}
}