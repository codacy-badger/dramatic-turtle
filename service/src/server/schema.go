package server

import (
	"errors"

	"../core"
	"github.com/graphql-go/graphql"
)

// Task struct
type Task struct {
	ID   string
	Name string
	Logs []LogEntry
}

// LogEntry struct
type LogEntry struct {
	ID    string
	Start string
	End   string
}

var (
	taskObject  *graphql.Object
	queryObject *graphql.Object
)

func init() {
	logEntryObject := graphql.NewObject(graphql.ObjectConfig{
		Name: "logEntry",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"start": &graphql.Field{
				Type: graphql.String,
			},
			"end": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	taskObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "task",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"logs": &graphql.Field{
				Type: graphql.NewList(logEntryObject),
			},
		},
	})
}

func getSchema(
	getTaskByID func(id string) Task,
	getAllTasks func() []Task) graphql.Schema {
	queryObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"tasks": &graphql.Field{
				Type: graphql.NewList(taskObject),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return getAllTasks(), nil
				},
			},
			"task": &graphql.Field{
				Type: taskObject,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(string)
					if ok {
						return getTaskByID(id), nil
					}
					return nil, errors.New("Could not find a matching task")
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryObject,
	})
	core.CheckErr(err)
	return schema
}
