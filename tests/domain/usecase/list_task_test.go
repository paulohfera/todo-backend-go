package usecase_test

import (
	"context"
	"testing"

	"github.com/paulohfera/todo-backend-go/configs"
	"github.com/paulohfera/todo-backend-go/internal/data/repository"
	"github.com/paulohfera/todo-backend-go/internal/domain/usecase"
	db "github.com/paulohfera/todo-backend-go/pkg/postgres"
)

func TestListTaskUseCase(t *testing.T) {
	t.Run("When call use case get a slice of tasks", func(t *testing.T) {
		ctx := context.Background()
		config := configs.GetConfigurations()
		conn := db.NewOrGetSingleton(config)
		taskReposytory := repository.NewTaskReposytory(conn)
		usecase := usecase.NewTaskUseCase(taskReposytory)
		tasks, err := usecase.ListTasks(ctx)
		if err != nil {
			t.Errorf("Error getting task list.")
		}

		if tasks[0].ID == 0 {
			t.Errorf("Error getting task list.")
		}
	})

}
