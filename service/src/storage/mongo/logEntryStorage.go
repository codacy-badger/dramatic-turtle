package mongo

import (
	"context"

	"../../core"
	"../models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// LogEntryStorage struct
type LogEntryStorage struct {
	mongo  *Mongo
	coll   *mongo.Collection
	taskID string
}

func createLogEntryStorage(m *Mongo, coll *mongo.Collection, taskID string) *LogEntryStorage {
	return &LogEntryStorage{
		mongo:  m,
		coll:   coll,
		taskID: taskID,
	}
}

// AppendLogEntry func
func (les *LogEntryStorage) AppendLogEntry(e *models.LogEntry) string {
	e.Start = e.Start.UTC()
	e.End = e.Start.UTC()

	e.ID = objectid.New().Hex()
	filter := bson.D{
		{"id", les.taskID},
	}
	update := bson.D{
		{"$push", bson.D{{"logs", *e}}},
	}
	_, err := les.coll.UpdateOne(context.Background(), filter, update)
	core.CheckErr(err)

	return e.ID
}

// ReadLogEntry func
func (les *LogEntryStorage) ReadLogEntry(id string) *models.LogEntry {
	task := les.mongo.GetTaskStorage().ReadTask(les.taskID)

	for _, le := range task.Logs {
		if id == le.ID {
			return &le
		}
	}
	return &models.LogEntry{}
}

// ReadLogEntries func
func (les *LogEntryStorage) ReadLogEntries(checkFunc func(id string) bool) []*models.LogEntry {
	task := les.mongo.GetTaskStorage().ReadTask(les.taskID)
	entries := []*models.LogEntry{}
	for i, le := range task.Logs {
		if checkFunc(le.ID) {
			entries = append(entries, &task.Logs[i])
		}
	}
	return entries
}
