package svc

import (
	"github.com/google/wire"
	"github.com/prf16/go-zero-box-rpc/app/internal/config"
	"github.com/prf16/go-zero-box-rpc/app/internal/model"
	"github.com/prf16/go-zero-box-rpc/app/internal/services"
	"github.com/prf16/go-zero-box-rpc/pkg"
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
