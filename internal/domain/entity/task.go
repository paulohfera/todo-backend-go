package entity

import (
	"time"
)

type Task struct {
	ID            int
	Title         string
	Description   string
	Due           *time.Time
	Done          bool
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	Notifications []string
}

func (i *Task) IsValid() bool {
	if i.Title == "" {

		i.Notifications = append(i.Notifications, "Title cannot be empty")
	}

	if i.Description == "" {
		i.Notifications = append(i.Notifications, "Description cannot be empty")
	}

	return len(i.Notifications) == 0
}

func NewTask(title string, description string, due *time.Time) *Task {
	return &Task{
		ID:            0,
		Title:         title,
		Description:   description,
		Due:           due,
		Done:          false,
		CreatedAt:     time.Now(),
		UpdatedAt:     nil,
		Notifications: []string{},
	}
}
