package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Todo struct {
	ID      int
	Task    string
	Duedate time.Time
	Urgency int
	Notes   string
	Created time.Time
}
