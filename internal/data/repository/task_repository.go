package repository

import (
	"context"
	"fmt"

	"github.com/paulohfera/todo-backend-go/internal/domain/entity"
	db "github.com/paulohfera/todo-backend-go/pkg/postgres"
)

type TaskReposytory struct {
	*db.DbContext
}

func NewTaskReposytory(db *db.DbContext) *TaskReposytory {
	return &TaskReposytory{db}
}

func (repository *TaskReposytory) Get(ctx context.Context, id int) (entity.Task, error) {
	sql := fmt.Sprintf("select * from task where id = %v", id)
	var task entity.Task
	row := repository.Pool.QueryRow(ctx, sql)
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Due, &task.Done, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return entity.Task{}, fmt.Errorf("TaskReposytory - Get - rows.Scan: %w", err)
	}

	return task, nil
}

func (repository *TaskReposytory) List(ctx context.Context) ([]entity.Task, error) {
	sql := "select * from task where done = false"
	rows, err := repository.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TaskReposytory - List - repository.Pool.Query: %w", err)
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Due, &task.Done, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("TaskReposytory - List - rows.Scan: %w", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *TaskReposytory) Add(ctx context.Context, item entity.Task) error {
	sql := `insert into task (title, description, due, done, createdat)
			values ($1, $2, $3, false, current_timestamp);`
	_, err := repository.Pool.Exec(ctx, sql, item.Title, item.Description, item.Due)
	if err != nil {
		return fmt.Errorf("TaskReposytory - Add - repository.Pool.Exec: %w", err)
	}

	return nil
}

func (repository *TaskReposytory) Update(ctx context.Context, item entity.Task) error {
	sql := `update task
			set title = $1,
			description = $2,
			due = $3,
			done = $4,
			updatedat = current_timestamp
			where id = $5;`
	_, err := repository.Pool.Exec(ctx, sql, item.Title, item.Description, item.Due, item.Done, item.ID)
	if err != nil {
		return fmt.Errorf("TaskReposytory - Update - repository.Pool.Exec: %w", err)
	}

	return nil
}

func (repository *TaskReposytory) Delete(ctx context.Context, id int) error {
	sql := `delete from task where id = $1;`
	_, err := repository.Pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("TaskReposytory - Delete - repository.Pool.Exec: %w", err)
	}

	return nil
}

func (repository *TaskReposytory) Complete(ctx context.Context, id int) error {
	sql := `update task set done = true where id = $1;`
	_, err := repository.Pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("TaskReposytory - Complete - repository.Pool.Exec: %w", err)
	}

	return nil
}
