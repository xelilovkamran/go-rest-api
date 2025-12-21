package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID	  int
}

func (e *Event) Save() error {
	res, err := db.DB.Exec(`INSERT INTO events (name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`,
		e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	
	e.ID = int(id)
	return nil
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query(`SELECT id, name, description, location, datetime, user_id FROM events`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID); err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func GetEventByID(id int) (*Event, error) {
	var event Event
	row := db.DB.QueryRow(`SELECT id, name, description, location, datetime, user_id FROM events WHERE id = ?`, id)
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) UpdateEvent() error {
	_, err := db.DB.Exec(`UPDATE events SET name = ?, description = ?, location = ?, datetime = ? WHERE id = ?`,
		e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Event) DeleteEvent() error {
	_, err := db.DB.Exec(`DELETE FROM events WHERE id = ?`, e.ID)
	return err
}

func (e Event) Register(userID int) error {
	_, err := db.DB.Exec(`INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`, e.ID, userID)
	return err
}

func (e Event) CancelRegistration(userID int) error {
	_, err := db.DB.Exec(`DELETE FROM registrations WHERE event_id = ? AND user_id = ?`, e.ID, userID)
	return err
}