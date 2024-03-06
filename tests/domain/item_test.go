package domain

import (
	"time"

	"testing"

	"github.com/paulohfera/todo-backend-go/domain"
)

func TestItemShouldRetunNotValidWhenMandatoryParametersAreMissing(t *testing.T) {
	item := domain.NewItem("", "", time.Now())

	got := item.Validate()
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestItemShouldRetunValidWhenAllMandatoryParametersAreOk(t *testing.T) {
	item := domain.NewItem("title", "description", time.Now())

	got := item.Validate()
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
