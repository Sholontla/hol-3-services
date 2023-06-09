package bigquery_service

import (
	"context"
	"restservice/config"

	"cloud.google.com/go/bigquery"
)

type BigQueryClient struct {
	config config.Config
}

func (c *BigQueryClient) CreateBigQueryClient() (*bigquery.Client, error) {
	project_id, _, _ := c.config.BigQueryConfig()
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, project_id)
	if err != nil {
		return nil, err
	}
	return client, nil
}
