package entity_test

import (
	"testing"

	"github.com/paulohfera/todo-backend-go/internal/domain/entity"
)

func TestDomainEntityTask(t *testing.T) {
	t.Run("When mandatory parameters are missing task must retun not valid", func(t *testing.T) {
		item := entity.NewTask("", "", nil)

		got := item.Validate()
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("When all mandatory parameters are ok task must retun valid", func(t *testing.T) {
		item := entity.NewTask("title", "description", nil)

		got := item.Validate()
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})
}
