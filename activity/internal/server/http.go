package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gitlab.okymo.com.tw/middle-end/kratos/pkg/conf/paladin"
	"gitlab.okymo.com.tw/middle-end/kratos/pkg/net/rpc/warden"
	pbactivity "test/activity"
	"time"
	"x-server/activity/internal/service"
)

// NewHTTPServer
// New new a http server.
func NewHTTPServer(svc *service.Service) (svr *http.Server, err error) {
	cfgs := &warden.ServerConfig{}
	ct := paladin.TOML{}
	if err = paladin.Get("http.txt").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfgs); err != nil {
		return
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if cfgs.Network != "" {
		opts = append(opts, http.Network(cfgs.Network))
	}
	if cfgs.Addr != "" {
		opts = append(opts, http.Address(cfgs.Addr))
	}
	if cfgs.Timeout != 0 {
		opts = append(opts, http.Timeout(time.Duration(cfgs.Timeout)))
	}
	svr = http.NewServer(opts...)
	pbactivity.RegisterActivityHTTPServer(svr, svc)
	return
}
