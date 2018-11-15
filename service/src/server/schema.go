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

//TODO (HAW): Resolve these functions into interface or sth.
func getSchema(
	getTaskByID func(id string) Task,
	getAllTasks func() []Task,
	createTask func(name string) string,
	createLogEntry func(taskID string) string) graphql.Schema {
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
			"createTask": &graphql.Field{
				Type:        graphql.String,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					name, ok := params.Args["name"].(string)
					if ok {
						return createTask(name), nil
					}
					return nil, errors.New("Could not create task. Please provide sufficient information.")
				},
			},
			"createLogEntry": &graphql.Field{
				Type:        graphql.String,
				Description: "",
				Args: graphql.FieldConfigArgument{
					"taskID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					taskID, ok := params.Args["taskID"].(string)
					if ok {
						return createLogEntry(taskID), nil
					}
					return nil, errors.New("Could not create logentry. Please provide sufficient information.")
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
