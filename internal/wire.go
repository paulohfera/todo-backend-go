//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package internal

import (
	"github.com/google/wire"
	"github.com/paulohfera/todo-backend-go/configs"
	"github.com/paulohfera/todo-backend-go/internal/data/repository"
	i "github.com/paulohfera/todo-backend-go/internal/domain/interface/repository"
	"github.com/paulohfera/todo-backend-go/internal/domain/usecase"
	"github.com/paulohfera/todo-backend-go/internal/rest"
	"github.com/paulohfera/todo-backend-go/internal/rest/handler"
	"github.com/paulohfera/todo-backend-go/pkg/api"
	db "github.com/paulohfera/todo-backend-go/pkg/postgres"
)

var deps = []interface{}{}

var providerSet wire.ProviderSet = wire.NewSet(
	db.NewOrGetSingleton,
	repository.NewTaskReposytory,
	usecase.NewTaskUseCase,
	handler.NewTaskHandler,
	rest.NewRestRouters,
	api.New,
	wire.Bind(new(i.ITaskRepository), new(*repository.TaskReposytory)),
)

func RegisterServices() *api.Api {
	wire.Build(providerSet, configs.GetConfigurations)
	return &api.Api{}
}

func RegisterServicesUseCase() *usecase.TaskUseCase {
	wire.Build(providerSet, configs.GetConfigurations)
	return &usecase.TaskUseCase{}
}
