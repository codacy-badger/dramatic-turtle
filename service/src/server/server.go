package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../config"
	"../core"
	"../storage"
	"../storage/mongo"
	"github.com/graphql-go/graphql"
)

var (
	dataStorage storage.IStorage
)

func stub(w http.ResponseWriter, r *http.Request) {
}

func getAllTasks() []Task {
	tasks := dataStorage.GetTaskStorage().ReadTasks(
		func(id string) bool {
			return true
		},
	)

	convTasks := []Task{}
	for _, e := range tasks {
		convTasks = append(convTasks, convertContractTask(e))
	}

	return convTasks
}

func getTaskByID(id string) Task {
	task := dataStorage.GetTaskStorage().ReadTask(id)
	return convertContractTask(task)
}

func endpoint(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	core.CheckErr(err)
	query := req["query"].(string)

	schema := getSchema(getTaskByID, getAllTasks)
	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	result := graphql.Do(params)
	payload, err := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// Start func
func Start(conf config.Config) {
	mongoConf := conf.Storage.MongoDB
	m := &mongo.Mongo{}
	m.Connect(mongoConf.Connection.URL, mongoConf.Connection.Database)
	dataStorage = m

	router := mux.NewRouter()

	router.HandleFunc("/v1/graphql", endpoint).Methods("GET")
	http.ListenAndServe(":"+strconv.Itoa(conf.Server.Port), router)
}
