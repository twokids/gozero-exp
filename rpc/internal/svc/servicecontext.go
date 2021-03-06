package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"zeroService/api/model"
	"zeroService/rpc/internal/config"
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
