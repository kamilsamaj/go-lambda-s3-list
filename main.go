package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"log"
	"os"
)

func getAWSInfo(event interface{}) {
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
}

func handleRequest(ctx context.Context, event events.CloudWatchEvent) (err error) {
	// log details about AWS environment
	getAWSInfo(event)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	} else {
		log.Println("AWS Config successfully loaded")
	}

	// Using the Config value, create the DynamoDB client
	s3Client := s3.NewFromConfig(cfg)

	lbOutput, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, bucket := range lbOutput.Buckets {
		log.Println(*bucket.Name)
	}
	return nil
}

func main() {
	lambda.Start(handleRequest)
}
