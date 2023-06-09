package config

type ConfigNodes struct {
	HTTP           HTTPServer     `mapstructure:"http_server"`
	BigQueryConfig BigQueryConfig `mapstructure:"bigquery_server"`
}

type HTTPServer struct {
	Port string `mapstructure:"port"`
}

type BigQueryConfig struct {
	ProjectID string `mapstructure:"project_id"`
	DatasetID string `mapstructure:"dataset_id"`
	TableID   string `mapstructure:"table_id"`
}
