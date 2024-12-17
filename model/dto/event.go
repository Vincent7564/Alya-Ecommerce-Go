package dto

import "time"

type AddEventRequest struct {
	EventName        string    `json:"event_name" validate:"required"`
	EventDescription string    `json:"event_description" validate:"required"`
	EventStartAt     time.Time `json:"event_start_at" validate:"required"`
	EventEndAt       time.Time `json:"event_end_at" validate:"required"`
	IsActive         bool      `json:"is_active" validate:"omitempty,required"`
}
