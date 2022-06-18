package mysql

import (
	"database/sql"
	"time"

	"kailashgautham.com/todotasks/pkg/models"
)

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) Insert(urgency int, duedate time.Time, task, notes string) (int, error) {
	result, err := m.DB.Exec(`INSERT INTO todolist (urgency, duedate, task, notes, created) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP())`, urgency, duedate, task, notes)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TodoModel) Complete(id int) error {

	return nil

}

func (m *TodoModel) OpenDetails(id int) (*models.Todo, error) {

	row := m.DB.QueryRow("SELECT id, task, duedate, urgency, notes, created FROM todolist where id = ?", id)

	s := &models.Todo{}
	err := row.Scan(&s.ID, &s.Task, &s.Duedate, &s.Urgency, &s.Notes, &s.Created)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil

}

func (m *TodoModel) ShowPending() ([]*models.Todo, error) {

	return nil, nil
}

func (m *TodoModel) ShowCompleted() {
}
