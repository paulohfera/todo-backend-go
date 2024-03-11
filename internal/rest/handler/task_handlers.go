package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/paulohfera/todo-backend-go/internal/domain/usecase"
)

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskUseCase.ListTasks(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
