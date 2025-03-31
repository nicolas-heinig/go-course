package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, datetime, user_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.Datetime, e.UserID).Scan(&e.ID)

	if err != nil {
		return err
	}

	return err
}

func (e Event) Update() error {
	query := `
		UPDATE events
		SET
			name = $1,
			description = $2,
			location = $3,
			datetime = $4
		WHERE id = $5
	`

	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.Datetime, e.ID)

	return err
}

func (e Event) Delete() error {
	query := `
		DELETE FROM events
		WHERE id = $1
	`

	_, err := db.DB.Exec(query, e.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES ($1, $2)`

	_, err := db.DB.Exec(query, e.ID, userId)

	return err
}

func (e Event) Cancel(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = $1 AND user_id = $2`

	_, err := db.DB.Exec(query, e.ID, userId)

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.Datetime,
			&event.UserID,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.Datetime,
		&event.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
