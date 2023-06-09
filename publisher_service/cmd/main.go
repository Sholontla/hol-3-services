package main

import "pubservice/internal/service/pubsub"

func main() {
	pub := pubsub.PublisherService{}
	pub.PubServer()
}
