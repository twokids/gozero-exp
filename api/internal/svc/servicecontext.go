package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"zeroService/api/internal/config"
	"zeroService/api/model"
	"zeroService/rpc/userserviceclient"
)

type ServiceContext struct {
	Config         config.Config
	Model          model.UserinfoModel           //新加
	UserServiceRpc userserviceclient.UserService //新加
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Model:          model.NewUserinfoModel(sqlx.NewMysql(c.DataSource)),         //新加
		UserServiceRpc: userserviceclient.NewUserService(zrpc.MustNewClient(c.Rpc)), //新加
	}
}
