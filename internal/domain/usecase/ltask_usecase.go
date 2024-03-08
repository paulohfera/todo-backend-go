package usecase

import (
	"context"

	"github.com/paulohfera/todo-backend-go/internal/domain/entity"
	"github.com/paulohfera/todo-backend-go/internal/domain/interface/repository"
)

type TaskUseCase struct {
	repository repository.ITaskRepository
}

func NewTaskUseCase(repository repository.ITaskRepository) *TaskUseCase {
	return &TaskUseCase{repository: repository}
}

func (u *TaskUseCase) ListTasks(ctx context.Context) ([]entity.Task, error) {
	tasks, err := u.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
