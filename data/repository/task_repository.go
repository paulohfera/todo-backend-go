package repository

import (
	"context"
	"fmt"

	"github.com/paulohfera/todo-backend-go/data/db"
	"github.com/paulohfera/todo-backend-go/data/entitymap"
	"github.com/paulohfera/todo-backend-go/domain/entity"
)

type TaskReposytory struct {
	*db.DbContext
}

func NewTaskReposytory(db *db.DbContext) *TaskReposytory {
	return &TaskReposytory{db}
}

func (repository *TaskReposytory) Get(ctx context.Context, ID int) (*entity.Task, error) {
	defer repository.Close()

	sql := fmt.Sprintf("select * from task where id = %v", ID)
	var task entitymap.TaskMap
	row := repository.Pool.QueryRow(ctx, sql)
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Due, &task.Done, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("TaskReposytory - Get - rows.Scan: %w", err)
	}

	return task.ToEntity(), nil
}
