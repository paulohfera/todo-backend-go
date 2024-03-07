package db_test

import (
	"testing"

	"github.com/paulohfera/todo-backend-go/configuration"
	"github.com/paulohfera/todo-backend-go/data/db"
)

func TestDataDbDbcontext(t *testing.T) {
	t.Run("When database connect don't get error", func(t *testing.T) {
		config := configuration.GetConfigurations()
		conn := db.NewOrGetSingleton(config)
		defer conn.Close()

		if conn == nil {
			t.Errorf("Error connecting to database.")
		}
	})
}
