package rest

import (
	"net/http"

	"github.com/paulohfera/todo-backend-go/internal/rest/handler"
)

type Rest struct {
	taskHandler *handler.TaskHandler
}

func NewRestRouters(taskHandler *handler.TaskHandler) *Rest {
	return &Rest{
		taskHandler: taskHandler,
	}
}

func (r *Rest) RegisterRouters(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/task/v1", r.taskHandler.List)
}
