package models

import (
	"database/sql"
	"fmt"
	"time"

	"kailashgautham.com/todotasks/pkg/models"
)

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) Insert(urgency int, duedate time.Time, task, notes string) (int, error) {
	result, err := m.DB.Exec("INSERT INTO todolist (urgency, duedate, task, notes, created) VALUES (?, ?, ?, ?, TIMESTAMP())", urgency, duedate, task, notes)
	if err != nil {
		return 0, fmt.Errorf("Insert: %v", err)
	}
}

func (m *TodoModel) Complete(id int) error {
	return nil
}

func (m *TodoModel) OpenDetails(id int) (*models.Todo, error) {
	return nil, nil
}

func (m *TodoModel) ShowPending() ([]*models.Todo, error) {
	return nil, nil
}

func (m *TodoModel) ShowCompleted() {
}
