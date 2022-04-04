package svc

import (
	"github.com/dfang/shorturl/api/internal/config"
	"github.com/dfang/shorturl/rpc/client"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer client.TransformService // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: client.NewTransformService(zrpc.MustNewClient(c.Transform)), // 手动代码
	}
}
