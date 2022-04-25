package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type Response struct {
	Message string `json:"message:"`
}

func handler(ctx context.Context, event events.S3Event) (Response, error) {
	log.Printf(`Recieved s3 event %+v`, event)
	return Response{
		Message: fmt.Sprintf("%+v event received", event),
	}, nil
}
func main() {
	lambda.Start(handler)
}
