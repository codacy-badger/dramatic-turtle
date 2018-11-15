package models

import (
	"time"
)

// LogEntry struct
type LogEntry struct {
	ID    string    `json:"id"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
