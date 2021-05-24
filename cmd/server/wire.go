// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/XMUMY/lost_found/internal/biz"
	"github.com/XMUMY/lost_found/internal/conf"
	"github.com/XMUMY/lost_found/internal/data"
	"github.com/XMUMY/lost_found/internal/server"
	"github.com/XMUMY/lost_found/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
