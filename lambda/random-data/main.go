package main

import (
	"encoding/json"
	"math/rand"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Data []int `json:"data"`
}

func handleEvent(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	s := make([]int, 100)

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(11)
	}

	sm, err := json.Marshal(&response{
		Data: s,
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
		Body: string(sm),
	}, nil
}

func main() {
	lambda.Start(handleEvent)
}