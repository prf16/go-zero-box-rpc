package svc

import (
	"github.com/prf16/go-zero-box-rpc/app/internal/config"
	"github.com/prf16/go-zero-box-rpc/app/internal/svc/command"
	"github.com/prf16/go-zero-box-rpc/app/internal/svc/model"
	"github.com/prf16/go-zero-box-rpc/app/internal/svc/queue"
	"github.com/prf16/go-zero-box-rpc/app/internal/svc/services"
	"github.com/prf16/go-zero-box-rpc/pkg"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewServiceContext,
	command.Provider,
	model.Provider,
	queue.Provider,
	services.Provider,
)

type ServiceContext struct {
	Config *config.Config
	Pkg    *pkg.Pkg

	Command  *command.Command
	Model    *model.Model
	Queue    *queue.Queue
	Services *services.Services
}

func NewServiceContext(config *config.Config, pkg *pkg.Pkg, command *command.Command, model *model.Model, queue *queue.Queue, services *services.Services) *ServiceContext {
	return &ServiceContext{Config: config, Pkg: pkg, Command: command, Model: model, Queue: queue, Services: services}
}
