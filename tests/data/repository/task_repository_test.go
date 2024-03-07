package repository_test

import (
	"context"
	"testing"

	"github.com/paulohfera/todo-backend-go/configuration"
	"github.com/paulohfera/todo-backend-go/data/db"
	"github.com/paulohfera/todo-backend-go/data/repository"
)

func TestTaskRepository(t *testing.T) {

	t.Run("When get task 1 get one task", func(t *testing.T) {
		ctx := context.Background()
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)
		defer conn.Close()

		taskReposytory := repository.NewTaskReposytory(conn)
		_, err := taskReposytory.Get(ctx, 1)
		if err != nil {
			t.Errorf("Error getting task 1.")
		}
	})

	t.Run("When get connection again get the same connection", func(t *testing.T) {
		config := configuration.GetConfigurations()
		want := db.NewOrGetSingleton(config)
		got := db.NewOrGetSingleton(config)
		defer got.Close()
		defer want.Close()

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
