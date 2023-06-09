package main

import "subservice/internal/service/pubsub"

func main() {
	sub := pubsub.SubsriberService{}
	sub.RunSubscriber()
}
