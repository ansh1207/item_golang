package utils

import (
	"encoding/json"
	"os"
)

type Options struct {
	UseNewUrlParser    bool `json:"useNewUrlParser"`
	UseUnifiedTopology bool `json:"useUnifiedTopology"`
}

type Config struct {
	Aerospike struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"aerospike"`
	Abapi struct {
		BaseUrl string `json:"baseUrl"`
	} `json:"abapi"`
	Campaigns struct {
		BaseUrl string `json:"baseUrl"`
	} `json:"campaigns"`
	Mongo struct {
		Url     string  `json:"url"`
		Options Options `json:"options"`
	} `json:"mongo"`
	Host   string `json:"host"`
	Port   string `json:"port"`
	DbType int    `json:"dbType"`
	DbUrl  string `json:"dbUrl"`
	DbName string `json:"dbName"`
}

func LoadConfiguration() (Config, error) {
	var config Config
	filename := "./config/" + os.Getenv("GO_ENV") + ".json"
	configfile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer configfile.Close()
	jsonPaser := json.NewDecoder(configfile)
	err = jsonPaser.Decode(&config)
	return config, err
}
