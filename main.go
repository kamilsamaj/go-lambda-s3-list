package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

func handleRequest(ctx context.Context, event events.SQSEvent) (err error) {
	// indent the JSON event for easy reading
	eventJson, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		log.Fatalln("Error when parsing event: ", event)
	}
	log.Printf("Event: %s", eventJson)
	// print all environment variables
	log.Println("ALL ENV Variables:")
	for _, envVar := range os.Environ() {
		log.Println(envVar)
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
