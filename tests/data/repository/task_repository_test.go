package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/paulohfera/todo-backend-go/configuration"
	"github.com/paulohfera/todo-backend-go/data/db"
	"github.com/paulohfera/todo-backend-go/data/repository"
	"github.com/paulohfera/todo-backend-go/domain/entity"
)

func TestTaskRepository(t *testing.T) {

	t.Run("When get task 1 get one task", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)

		taskReposytory := repository.NewTaskReposytory(conn)
		_, err := taskReposytory.Get(ctx, 1)
		if err != nil {
			t.Errorf("Error getting task 1.")
		}
	})

	t.Run("When get task 0 no rows should return", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)

		taskReposytory := repository.NewTaskReposytory(conn)
		_, err := taskReposytory.Get(ctx, 0)
		if err == nil {
			t.Errorf("Error getting task 0.")
		}
	})

	t.Run("When get task list get a slice of task", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)

		taskReposytory := repository.NewTaskReposytory(conn)
		tasks, err := taskReposytory.List(ctx)
		if err != nil {
			t.Errorf("Error getting task list.")
		}

		if tasks[0].ID == 0 {
			t.Errorf("Error getting task list.")
		}
	})

	t.Run("When insert valid task should not retunr error", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)

		taskReposytory := repository.NewTaskReposytory(conn)
		due := time.Now().AddDate(0, 1, 0)
		task := entity.NewTask("task unity test", "unit test", &due)
		err := taskReposytory.Add(ctx, *task)
		if err != nil {
			t.Errorf("Error adding new valid task.")
		}
	})

	t.Run("When update valid task should not retunr error", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)

		taskReposytory := repository.NewTaskReposytory(conn)
		task, _ := taskReposytory.Get(ctx, 1)
		task.Title = "updated title"
		task.Description = "updated description"
		due := time.Now().AddDate(0, 1, 0)
		task.Due = &due

		err := taskReposytory.Update(ctx, task)
		if err != nil {
			t.Errorf("Error updating task.")
		}
	})

	t.Run("When id is valid should delete and not retunr error", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)

		taskReposytory := repository.NewTaskReposytory(conn)
		err := taskReposytory.Delete(ctx, 2)
		if err != nil {
			t.Errorf("Error deleting a task.")
		}
	})

	t.Run("When id is valid should complete and not retunr error", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)

		taskReposytory := repository.NewTaskReposytory(conn)
		err := taskReposytory.Complete(ctx, 3)
		if err != nil {
			t.Errorf("Error completing a task.")
		}
	})

	t.Run("When get connection again get the same connection", func(t *testing.T) {
		config := configuration.GetConfigurations()
		want := db.NewOrGetSingleton(config)
		got := db.NewOrGetSingleton(config)

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
