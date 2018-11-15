package config

import (
	"encoding/json"
	"io/ioutil"

	"../core"
)

// Config struct
type Config struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Storage struct {
		MongoDB struct {
			Active     bool `json:"active"`
			Connection struct {
				URL      string `json:"url"`
				Database string `json:"database"`
			} `json:"connection"`
		} `json:"mongodb"`
	} `json:"storage"`
}

// LoadConfig func
func LoadConfig(fileName string) Config {
	var data Config
	file, err := ioutil.ReadFile(fileName)
	core.CheckErr(err)
	err = json.Unmarshal(file, &data)
	core.CheckErr(err)

	return data
}
