package server

import (
	"../storage/models"
)

func convertContractLogEntry(logEntry *models.LogEntry) LogEntry {
	return LogEntry{
		ID:    logEntry.ID,
		Start: logEntry.Start.UTC().String(),
		End:   logEntry.End.UTC().String(),
	}
}

func convertContractTask(task *models.Task) Task {
	cTask := Task{
		ID:   task.ID,
		Name: task.Name,
		Logs: []LogEntry{},
	}
	for _, e := range task.Logs {
		cTask.Logs = append(cTask.Logs, convertContractLogEntry(&e))
	}
	return cTask
}
