//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/prf16/go-zero-box-rpc/app/internal/config"
	"github.com/prf16/go-zero-box-rpc/app/internal/svc"
	"github.com/prf16/go-zero-box-rpc/pkg"
)

type App struct {
	config *config.Config
	svcCtx *svc.ServiceContext
	pkg    *pkg.Pkg
}

func NewApp(config *config.Config, svcCtx *svc.ServiceContext, pkg *pkg.Pkg) *App {
	return &App{config: config, svcCtx: svcCtx, pkg: pkg}
}

func initApp(c *config.Config) *App {
	wire.Build(
		config.Provider,
		svc.Provider,
		pkg.Provider,
		NewApp,
	)
	return nil
}
