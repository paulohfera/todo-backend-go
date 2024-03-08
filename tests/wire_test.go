package internal_test

import (
	"context"
	"testing"

	"github.com/paulohfera/todo-backend-go/internal"
)

func TestWire(t *testing.T) {
	t.Run("When wire loaded list task use case shoul return list", func(t *testing.T) {
		ctx := context.Background()
		w := internal.RegisterServicesUseCase()
		tasks, err := w.ListTasks(ctx)
		if err != nil {
			t.Errorf("Error getting task list.")
		}

		if tasks[0].ID == 0 {
			t.Errorf("Error getting task list.")
		}
	})
}
