package db_test

import (
	"testing"

	"github.com/paulohfera/todo-backend-go/configs"
	db "github.com/paulohfera/todo-backend-go/pkg/postgres"
)

func TestDataDbDbcontext(t *testing.T) {
	t.Run("When database connect don't get error", func(t *testing.T) {
		config := configs.GetConfigurations()
		conn := db.NewOrGetSingleton(config)
		defer conn.Close()

		if conn == nil {
			t.Errorf("Error connecting to database.")
		}
	})
}
