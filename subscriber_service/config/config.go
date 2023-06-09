package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
}

var Conf = JsonConfigNodes()
var vp *viper.Viper

func JsonConfigNodes() ConfigNodes {
	// Initialize Viper for read config files
	vp = viper.New()
	var config ConfigNodes

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("config_files")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		log.Println("Error while reading config.json")
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Println("Error Unmarrsall", err)
	}
	return config
}

// HTTP congiguration
func (c Config) HTTPConfig() string {
	var (
		data = JsonConfigNodes()
		port = data.HTTP.Port
	)
	return port
}

// MongoDB configuration
func (c Config) BigQueryConfig() (string, string, string, string) {
	var (
		data      = JsonConfigNodes()
		projectID = data.BigQueryConfig.ProjectID
		datasetID = data.BigQueryConfig.DatasetID
		tableID   = data.BigQueryConfig.TableID
		topicID   = data.BigQueryConfig.TopicID
	)
	return projectID, datasetID, tableID, topicID
}
