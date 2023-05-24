package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gitlab.okymo.com.tw/middle-end/kratos/pkg/conf/paladin"
	"gitlab.okymo.com.tw/middle-end/kratos/pkg/net/rpc/warden"
	pbactivity "test/activity"
	"time"
	"x-server/activity/internal/service"
)

// New new a grpc server.
func NewGRPCServer(svc *service.Service) (svr *grpc.Server, err error) {
	ct := paladin.TOML{}
	cfg := &warden.ServerConfig{}

	if err = paladin.Get("grpc.txt").Unmarshal(&ct); err != nil {
		cfg = nil
		return
	} else if err = ct.Get("Server").UnmarshalTOML(cfg); err != nil {
		cfg = nil
		return
	}

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if cfg.Network != "" {
		opts = append(opts, grpc.Network(cfg.Network))
	}
	if cfg.Addr != "" {
		opts = append(opts, grpc.Address(cfg.Addr))
	}
	if cfg.Timeout != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(cfg.Timeout)))
	}
	svr = grpc.NewServer(opts...)
	pbactivity.RegisterActivityServer(svr, svc)
	return
}
