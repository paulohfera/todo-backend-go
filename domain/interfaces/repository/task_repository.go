package repository

import (
	"context"

	"github.com/paulohfera/todo-backend-go/domain/entity"
)

type ITaskRepository interface {
	Get(ctx context.Context, ID int) (entity.Task, error)
	List(ctx context.Context) ([]entity.Task, error)
	Add(ctx context.Context, item entity.Task) error
	Update(ctx context.Context, item entity.Task) error
	Delete(ctx context.Context, ID int) error
	Complete(ctx context.Context, ID int) error
}
