package domain

import (
	"time"
)

type Item struct {
	Id            int
	Title         string
	Description   string
	Due           time.Time
	Done          bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Notifications []string
}

func (i *Item) Validate() bool {
	if i.Title == "" {
		i.Notifications = append(i.Notifications, "Title cannot be empty")
	}

	if i.Description == "" {
		i.Notifications = append(i.Notifications, "Description cannot be empty")
	}

	return len(i.Notifications) == 0
}

func NewItem(title string, description string, due time.Time) *Item {
	return &Item{
		Title:       title,
		Description: description,
		Due:         due,
		Done:        false,
		CreatedAt:   time.Now(),
	}
}
