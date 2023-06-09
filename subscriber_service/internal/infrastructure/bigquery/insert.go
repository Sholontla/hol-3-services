package bigquery

import (
	"context"

	"subservice/internal/entity"

	"cloud.google.com/go/bigquery"
)

func InsertRow(client *bigquery.Client, data entity.MyData, datasetID, tableID string) error {
	// Define the field schemas for each column
	schema := bigquery.Schema{
		{Name: "insertID", Type: bigquery.StringFieldType},
		{Name: "column1", Type: bigquery.StringFieldType},
		{Name: "column2", Type: bigquery.IntegerFieldType},
		{Name: "column3", Type: bigquery.FloatFieldType},
		{Name: "column4", Type: bigquery.StringFieldType},
		{Name: "created_at", Type: bigquery.StringFieldType},
	}

	// Create the BigQuery row object with the schema
	row := &bigquery.StructSaver{
		Schema:   schema,
		InsertID: data.InsertID,
		Struct:   data,
	}

	// Make the BigQuery API request to insert the row
	err := client.Dataset(datasetID).Table(tableID).Inserter().Put(context.Background(), []*bigquery.StructSaver{row})
	if err != nil {
		return err
	}

	return nil
}
