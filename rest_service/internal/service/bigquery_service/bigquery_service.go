package bigquery_service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restservice/config"
	"time"

	"restservice/internal/entity"
	insert "restservice/internal/infrastructure/bigquery"
	"restservice/internal/service/table_creation"

	"cloud.google.com/go/bigquery"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type BigQueryService struct {
	config config.Config
	client BigQueryClient
	table  table_creation.CreateDataSetService
}

func (conf *BigQueryService) insertRowHandler(w http.ResponseWriter, r *http.Request) {
	_, dataset_id, table_id := conf.config.BigQueryConfig()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	client, err := conf.client.CreateBigQueryClient()
	if err != nil {
		log.Printf("Error creating BigQuery client: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Extract the data from the request
	var data entity.MyData
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	datasetID := dataset_id
	tableID := table_id

	// Create a new instance of entity.MyData and populate its fields
	rowData := entity.MyData{
		InsertID:  uuid.New().String(),
		Column1:   data.Column1,
		Column2:   data.Column2,
		Column3:   data.Column3,
		Column4:   data.Column4,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05"),
		// Add more fields as needed
	}
	// Call the insertRow function to insert the row
	err = insert.InsertRow(client, rowData, datasetID, tableID) // Pass &rowData as a pointer to entity.MyData
	if err != nil {
		log.Printf("Error inserting row into BigQuery: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Row inserted successfully")
}

// Rest of the code remains the same

func (conf *BigQueryService) readRowHandler(w http.ResponseWriter, r *http.Request) {

	_, dataset_id, table_id := conf.config.BigQueryConfig()

	client, err := conf.client.CreateBigQueryClient()
	if err != nil {
		log.Printf("Error creating BigQuery client: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set up the BigQuery API request
	datasetID := dataset_id
	tableID := table_id

	// Specify the row ID or any other unique identifier to read a specific row
	rowID := "your-row-id"

	// Make the BigQuery API request to get the row iterator
	it := client.Dataset(datasetID).Table(tableID).Read(context.Background())

	// Loop through the iterator until we find the desired row
	var row []bigquery.Value
	for {
		err := it.Next(&row)
		if err == iterator.Done {
			break // No more rows to read
		}
		if err != nil {
			log.Printf("Error reading row from BigQuery: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Check if the current row matches the desired row ID
		if row[0] == rowID {
			// Found the row, do something with it
			fmt.Fprintf(w, "Row: %v", row)
			return
		}
	}

	// If the loop completes without finding the row, it means the row ID was not found
	log.Printf("Row not found in BigQuery")
	http.Error(w, "Row not found", http.StatusNotFound)
}

func (serv *BigQueryService) Endpoints() {
	port := serv.config.HTTPConfig()

	http.HandleFunc("/createtable", serv.table.CreateTableHandler)
	http.HandleFunc("/insert", serv.insertRowHandler)
	http.HandleFunc("/read", serv.readRowHandler)

	log.Printf("Running rest service on port: %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
