package repository

import (
	"time"
)

type Event struct {
	EventID     string    `db:"event_id"`
	AggregateID string    `db:"aggregate_id"` // user or data id
	EventType   string    `db:"event_type"`
	EventData   string    `db:"event_data"` // JSONB type
	CreatedAt   time.Time `db:"created_at"`
}
