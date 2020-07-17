package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"strings"

	daprd "github.com/dapr/go-sdk/service/http"
)

var (
	logger    = log.New(os.Stdout, "", 0)
	address   = getEnvVar("ADDRESS", ":8080")
	topicName = getEnvVar("TOPIC_NAME", "events")
)

func main() {
	// create a Dapr service
	s := daprd.NewService()

	// add some topic subscriptions
	err := s.AddTopicEventHandler("events", "/events", eventHandler)
	if err != nil {
		logger.Fatalf("error adding topic subscription: %v", err)
	}

	// start the service
	if err = s.Start(address); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("error starting service: %v", err)
	}
}

func eventHandler(ctx context.Context, e daprd.TopicEvent) error {
	logger.Printf(
		"event - Source: %s, Topic:%s, ID:%s, DataContentType:%s",
		e.Source, e.Topic, e.ID, e.DataContentType,
	)

	// TODO: do something with the cloud event data
	logger.Println(e.Data)

	return nil
}

func getEnvVar(key, fallbackValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return strings.TrimSpace(val)
	}
	return fallbackValue
}
