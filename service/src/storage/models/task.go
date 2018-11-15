package models

// Task struct
type Task struct {
	ID   string     `json:"id"`
	Name string     `json:"name"`
	Logs []LogEntry `json:"logs"`
}
