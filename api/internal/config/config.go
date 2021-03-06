package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

//加入下面两个配置声明
type Config struct {
	rest.RestConf
	DataSource string             //新加
	Cache      cache.CacheConf    //新加
	Rpc        zrpc.RpcClientConf //新加
}
