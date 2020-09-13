package main

import (
	"encoding/json"
	"math/rand"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Data int `json:"data"`
}

func handleEvent(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	i, err := json.Marshal(&response{
		Data: rand.Intn(11),
	})

	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
		}, err
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(i),
	}, nil
}

func main() {
	lambda.Start(handleEvent)
}