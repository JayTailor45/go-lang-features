package models

import "jaytailor.com/event-management/db"

type Registration struct {
	ID        int64  `binding:"required"`
	EventID   int64  `binding:"required"`
	EventName string `binding:"required"`
	UserID    int64  `binding:"required"`
	UserEmail string `binding:"required"`
}

func (e *Registration) ListAllRegistrations() ([]Registration, error) {
	query := `
	SELECT r.id, r.user_id, r.event_id, u.email AS user_email, e.name AS event_name
	FROM registrations r
	JOIN events e ON r.event_id = e.id
	JOIN users u ON r.user_id = u.id;
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Scan()

	eventRegistrations := []Registration{}

	for rows.Next() {
		var reg Registration
		err := rows.Scan(&reg.ID, &reg.UserID, &reg.EventID, &reg.UserEmail, &reg.EventName)
		if err != nil {
			return nil, err
		}
		eventRegistrations = append(eventRegistrations, reg)
	}

	return eventRegistrations, err
}
