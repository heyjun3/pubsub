package main

import (
	"strconv"

	"github.com/heyjun3/pubsub"
)

func main() {
	projectID := "gsheet-355401"
	topicID := "order-queue"
	total := 50
	messages := make([]pubsub.Message, 0, total)
	for i := 0; i < total; i++ {
		messages = append(messages, pubsub.Message{
			Message: "message" + strconv.Itoa(i),
			OrderingKey: "key1",
		})
	}
	pubsub.PublishWithOrderingKey(projectID, topicID, messages)
}
