package entity_test

import (
	"testing"
	"time"

	"github.com/paulohfera/todo-backend-go/domain/entity"
)

func TestDomainEntityTask(t *testing.T) {
	t.Run("When mandatory parameters are missing task must retun not valid", func(t *testing.T) {
		item := entity.NewTask("", "", time.Now())

		got := item.Validate()
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("When all mandatory parameters are ok task must retun valid", func(t *testing.T) {
		item := entity.NewTask("title", "description", time.Now())

		got := item.Validate()
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})
}
