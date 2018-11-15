package mongo

import (
	"context"

	"../../core"
	"../../storage"
	"../models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// TaskStorage struct
type TaskStorage struct {
	mongo *Mongo
	coll  *mongo.Collection
}

func createTaskStorage(m *Mongo, id string) *TaskStorage {
	return &TaskStorage{mongo: m, coll: m.db.Collection(id)}
}

// StoreTask func
func (ts *TaskStorage) StoreTask(task *models.Task) string {
	id, err := ts.coll.InsertOne(context.Background(), task)
	strId := ts.mongo.getID(id)
	objId, err := objectid.FromHex(strId)
	core.CheckErr(err)

	filter := bson.D{
		{"_id", objId},
	}
	update := bson.D{
		{"$set", bson.D{{"id", strId}}},
	}
	_, err = ts.coll.UpdateOne(context.Background(), filter, update)

	core.CheckErr(err)
	return strId
}

// ReadTask func
func (ts *TaskStorage) ReadTask(id string) *models.Task {
	var res models.Task
	ts.coll.FindOne(context.Background(),
		bson.D{{"id", id}}).Decode(&res)
	return &res
}

// ReadTasks func
func (ts *TaskStorage) ReadTasks(checkFunc func(id string) bool) []*models.Task {
	ctx := context.Background()
	var res = []*models.Task{}
	cursor, err := ts.coll.Find(ctx, nil)
	core.CheckErr(err)
	defer cursor.Close(ctx)

	current := bson.D{}
	for cursor.Next(ctx) {
		cursor.Decode(current)
		id := current.Map()["id"].(string)
		var currentTask models.Task
		if checkFunc(id) {
			cursor.Decode(&currentTask)
			res = append(res, &currentTask)
		}
	}
	return res
}

// GetLog func
func (ts *TaskStorage) GetLog(id string) storage.ILogEntryStorage {
	return createLogEntryStorage(ts.mongo, ts.coll, id)
}
