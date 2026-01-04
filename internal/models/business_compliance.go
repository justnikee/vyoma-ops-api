package models

import (
	"time"

	"github.com/google/uuid"
)

type BusinessCompliance struct {
	BusinessID uuid.UUID `json:"business_id"`
	RuleID     uuid.UUID `json:"rule_id"`
	Status     string    `json:"status"` // pending | completed | overdue
	DueDate    time.Time `json:"due_date"`
}
