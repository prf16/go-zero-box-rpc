package svc

import (
	"github.com/google/wire"
	"go-zero-box-rpc/app/internal/config"
	"go-zero-box-rpc/app/internal/model"
	"go-zero-box-rpc/app/internal/services"
	"go-zero-box-rpc/pkg"
)

var Provider = wire.NewSet(
	NewServiceContext,
	model.Provider,
	services.Provider,
)

type ServiceContext struct {
	Config  *config.Config
	Model   *model.Model
	Service *services.Services
	Pkg     *pkg.Pkg
}

func NewServiceContext(config *config.Config, model *model.Model, service *services.Services, pkg *pkg.Pkg) *ServiceContext {
	return &ServiceContext{Config: config, Model: model, Service: service, Pkg: pkg}
}
