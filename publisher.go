package pubsub

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type Message struct {
	Message     string
	OrderingKey string
}

func PublishWithOrderingKey(projectID, topicID string, messages []Message) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID,
		option.WithEndpoint("asia-northeast1-pubsub.googleapis.com:443"),
	)
	if err != nil {
		return err
	}
	defer client.Close()

	var wg sync.WaitGroup
	var totalErrors uint64
	topic := client.Topic(topicID)
	// topic.EnableMessageOrdering = true

	for _, m := range messages {
		res := topic.Publish(ctx, &pubsub.Message{
			Data:        []byte(m.Message),
			// OrderingKey: m.OrderingKey,
		})
		wg.Add(1)
		go func(res *pubsub.PublishResult) {
			defer wg.Done()
			_, err := res.Get(ctx)
			if err != nil {
				slog.Error("Failed to publish: ", "error", err)
				atomic.AddUint64(&totalErrors, 1)
				return
			}
		}(res)
	}
	wg.Wait()

	if totalErrors > 0 {
		return fmt.Errorf("%d of %d messages did not publish successfully", totalErrors, len(messages))
	}
	slog.Info(fmt.Sprintf("Published %d messages with ordering keys successfully", len(messages)))
	return nil
}
