package repository

import (
	"context"
	"fmt"

	"github.com/paulohfera/todo-backend-go/data/db"
	"github.com/paulohfera/todo-backend-go/domain/entity"
)

type TaskReposytory struct {
	*db.DbContext
}

func NewTaskReposytory(db *db.DbContext) *TaskReposytory {
	return &TaskReposytory{db}
}

func (repository *TaskReposytory) Get(ctx context.Context, ID int) (entity.Task, error) {
	sql := fmt.Sprintf("select * from task where id = %v", ID)
	var task entity.Task
	row := repository.Pool.QueryRow(ctx, sql)
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Due, &task.Done, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return entity.Task{}, fmt.Errorf("TaskReposytory - Get - rows.Scan: %w", err)
	}

	return task, nil
}

func (repository *TaskReposytory) List(ctx context.Context) ([]entity.Task, error) {
	// defer repository.Close()

	sql := "select * from task where done = false"
	rows, err := repository.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TaskReposytory - List - repository.Pool.Query: %w", err)
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Due, &task.Done, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("TaskReposytory - List - rows.Scan: %w", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *TaskReposytory) Add(ctx context.Context, item entity.Task) error {
	sql := fmt.Sprintf(`insert into task (title, description, due, done, createdat)
	values ('%v', '%v', $1, false, current_timestamp);`, item.Title, item.Description)
	_, err := repository.Pool.Exec(ctx, sql, item.Due)
	if err != nil {
		return fmt.Errorf("TaskReposytory - Add - repository.Pool.Exec: %w", err)
	}

	return nil
}
