package entitymap

import (
	"database/sql"
	"time"

	"github.com/paulohfera/todo-backend-go/domain/entity"
)

type TaskMap struct {
	ID          int
	Title       string
	Description string
	Due         sql.NullTime
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

func (i *TaskMap) ToEntity() *entity.Task {
	return &entity.Task{
		ID:          i.ID,
		Title:       i.Title,
		Description: i.Description,
		Due: func() time.Time {
			if i.Due.Valid {
				return i.Due.Time
			}
			return time.Time{}
		}(),
		Done:      i.Done,
		CreatedAt: i.CreatedAt,
		UpdatedAt: func() time.Time {
			if i.UpdatedAt.Valid {
				return i.UpdatedAt.Time
			}
			return time.Time{}
		}(),
	}
}
