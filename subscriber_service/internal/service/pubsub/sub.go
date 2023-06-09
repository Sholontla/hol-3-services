package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"subservice/config"
	"subservice/internal/entity"
	insert "subservice/internal/infrastructure/bigquery"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/pubsub"
)

type SubsriberService struct {
	config config.Config
}

func (sub *SubsriberService) Subscriber(ctx context.Context, msg pubsub.Message) error {
	project_id, dataset_id, table_id, _ := sub.config.BigQueryConfig()
	// Extract the data from the Pub/Sub message
	var data entity.MyData
	err := json.Unmarshal(msg.Data, &data)
	if err != nil {
		log.Printf("Error decoding Pub/Sub message data: %v", err)
		return err
	}

	projectID := project_id
	// Set up the BigQuery API request
	datasetID := dataset_id
	tableID := table_id

	// Create a BigQuery client
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Error creating BigQuery client: %v", err)
		return err
	}

	// Call the InsertRow function to insert the row
	err = insert.InsertRow(client, data, datasetID, tableID)
	if err != nil {
		log.Printf("Error inserting row into BigQuery: %v", err)
		return err
	}

	log.Println("Row inserted successfully")

	return nil
}

func (sub *SubsriberService) StartSubscriber(ctx context.Context, projectID, subscriptionID string) error {
	// Create a Pub/Sub client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Error creating Pub/Sub client: %v", err)
		return err
	}

	// Create a Pub/Sub subscription
	subscription := client.Subscription(subscriptionID)

	// Receive and process messages from the subscription
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		err := sub.Subscriber(ctx, *msg)
		if err != nil {
			log.Fatal(err)
		}

		// Acknowledge the message to mark it as processed
		msg.Ack()
	})
	if err != nil {
		log.Printf("Error receiving Pub/Sub messages: %v", err)
		return err
	}

	return nil
}

func (sub *SubsriberService) RunSubscriber() {
	project_id, _, _, subscription_id := sub.config.BigQueryConfig()
	// Set up the Google Cloud project and subscription information
	projectID := project_id
	subscriptionID := subscription_id

	// Create a context
	ctx := context.Background()

	// Start the subscriber
	err := sub.StartSubscriber(ctx, projectID, subscriptionID)
	if err != nil {
		log.Fatal(err)
	}
}
