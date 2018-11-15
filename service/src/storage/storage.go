package storage

import (
	"./models"
)

// ITaskStorage interface
type ITaskStorage interface {
	StoreTask(t *models.Task) string
	ReadTask(id string) *models.Task
	ReadTasks(checkFunc func(id string) bool) []*models.Task

	GetLog(id string) ILogEntryStorage
}

// ILogEntryStorage interface
type ILogEntryStorage interface {
	AppendLogEntry(e *models.LogEntry) string
	ReadLogEntry(id string) *models.LogEntry
	ReadLogEntries(checkFunc func(id string) bool) []*models.LogEntry
}

// IStorage interface
type IStorage interface {
	GetTaskStorage() ITaskStorage
}
