package psw

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CreateSubscriptionIfNotExists will create a subscription if it doesnt exist, otherwise return the existing one
func CreateSubscriptionIfNotExists(client *pubsub.Client, subName string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	// Create a topic to subscribe to.
	s, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:             topic,
		AckDeadline:       5 * time.Minute,
		RetentionDuration: time.Hour * 24 * 2,
	})
	if err != nil && grpc.Code(err) != codes.AlreadyExists {
		return nil, fmt.Errorf("Failed to create the pubsub topic: %s", err.Error())
	} else if grpc.Code(err) == codes.AlreadyExists {
		return client.Subscription(subName), nil
	}
	return s, nil
}

// CreateTopicIfNotExists will create a topic if it doesnt exist, otherwise return the existing one
func CreateTopicIfNotExists(client *pubsub.Client, topicName string) (*pubsub.Topic, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	// Create a topic to subscribe to.
	t, err := client.CreateTopic(ctx, topicName)
	if err != nil && grpc.Code(err) != codes.AlreadyExists {
		return nil, fmt.Errorf("Failed to create the pubsub topic: %s", err.Error())
	} else if grpc.Code(err) == codes.AlreadyExists {
		return client.Topic(topicName), nil
	}
	return t, nil
}
