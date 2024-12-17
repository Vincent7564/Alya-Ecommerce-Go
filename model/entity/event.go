package entity

import "time"

type Event struct {
	EventName        string    `json:"event_name"`
	EventDescription string    `json:"event_description"`
	EventStartAt     time.Time `json:"event_start_at"`
	EventEndAt       time.Time `json:"event_end_at"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedAt        time.Time `json:"updated_at"`
	UpdatedBy        string    `json:"updated_by"`
}
