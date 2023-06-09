package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pubservice/config"
	"pubservice/internal/entity"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
)

type PublisherService struct {
	config config.Config
}

func Publisher(ctx context.Context, projectID, topicID string) {
	// Create a Pub/Sub client
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}

	// Get a topic reference
	topic := client.Topic(topicID)

	// Publish messages to the topic
	for i := 1; i <= 10; i++ {
		// Create the data to publish
		data := &entity.MyData{
			InsertID:  uuid.New().String(),
			Column1:   "Value 1",
			Column2:   i,
			Column3:   3.14,
			Column4:   "Value 4",
			CreatedAt: time.Now().Format(time.RFC3339),
		}

		// Convert data to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Failed to marshal JSON data: %v", err)
			continue
		}

		message := &pubsub.Message{
			Data: jsonData,
		}

		// Publish the message
		result := topic.Publish(ctx, message)

		// Get the server-generated message ID
		messageID, err := result.Get(ctx)
		if err != nil {
			log.Printf("Failed to publish message %d: %v", i, err)
		} else {
			log.Printf("Published message %d, Message ID: %s", i, messageID)
		}

		time.Sleep(1 * time.Second) // Add a delay between each message
	}

	// Close the client connection
	client.Close()
}

func (pub *PublisherService) PublishHandler(w http.ResponseWriter, r *http.Request) {
	project_id, _, _, topic_id := pub.config.BigQueryConfig()
	// Set up the Google Cloud project and topic information
	projectID := project_id
	topicID := topic_id

	// Create a context
	ctx := context.Background()

	// Publish messages
	Publisher(ctx, projectID, topicID)

	// Return a success message
	w.Write([]byte("Messages published successfully"))
}

func (pub *PublisherService) PubServer() {
	port := pub.config.HTTPConfig()
	http.HandleFunc("/publisher", pub.PublishHandler)
	log.Printf("Running publisher service on port: %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
