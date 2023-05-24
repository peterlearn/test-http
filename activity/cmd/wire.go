//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"x-server/activity/internal/server"
	"x-server/activity/internal/service"

	"github.com/google/wire"
)

//	func initApp() (*kratos.App, func(), error) {
//		panic(wire.Build(dao.ProviderSet, service.ProviderSet, coredao.ProviderSet, server.ProviderSet, newApp))
//	}
func initApp() (*kratos.App, func(), error) {
	panic(wire.Build(service.ProviderSet, server.ProviderSet, newApp))
}
