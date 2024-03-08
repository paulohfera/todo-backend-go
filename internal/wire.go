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
	db "github.com/paulohfera/todo-backend-go/pkg/postgres"
)

var providerSet wire.ProviderSet = wire.NewSet(
	db.NewOrGetSingleton,
	repository.NewTaskReposytory,

	usecase.NewTaskUseCase,
	wire.Bind(new(i.ITaskRepository), new(*repository.TaskReposytory)),
)

func RegisterServicesUseCase() *usecase.TaskUseCase {
	wire.Build(providerSet, configs.GetConfigurations)
	return &usecase.TaskUseCase{}
}
