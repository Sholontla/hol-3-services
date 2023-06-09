package table_creation

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restservice/config"

	"cloud.google.com/go/bigquery"
)

type CreateDataSetService struct {
	config config.Config
}

func createDataset(projectID, datasetID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	meta := &bigquery.DatasetMetadata{
		Name: datasetID,
	}

	if err := client.Dataset(datasetID).Create(ctx, meta); err != nil {
		return err
	}

	return nil
}

func (conf *CreateDataSetService) createTable(projectID, datasetID, tableID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	schema := bigquery.Schema{
		{
			Name: "insertID",
			Type: bigquery.StringFieldType,
		},
		{
			Name: "column1",
			Type: bigquery.StringFieldType,
		},
		{
			Name: "column2",
			Type: bigquery.IntegerFieldType,
		},
		{
			Name: "column3",
			Type: bigquery.FloatFieldType,
		},
		{
			Name: "column4",
			Type: bigquery.StringFieldType,
		},
		{
			Name: "created_at",
			Type: bigquery.StringFieldType,
		},
		// Add more fields as needed
	}

	table := &bigquery.TableMetadata{
		Name:   tableID,
		Schema: schema,
	}

	if err := client.Dataset(datasetID).Table(tableID).Create(ctx, table); err != nil {
		return err
	}

	return nil
}

func (conf *CreateDataSetService) CreateTable() {
	project_id, dataset_id, table_id := conf.config.BigQueryConfig()
	projectID := project_id
	datasetID := dataset_id
	tableID := table_id

	if err := createDataset(projectID, datasetID); err != nil {
		log.Fatal(err)
	}

	if err := conf.createTable(projectID, datasetID, tableID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dataset and table created successfully")
}

func (conf *CreateDataSetService) CreateTableHandler(w http.ResponseWriter, r *http.Request) {
	// Call the CreateTable function
	conf.CreateTable()

	// Write a response
	fmt.Fprint(w, "Table created successfully")
}
