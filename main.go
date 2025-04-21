package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

type MyEvent struct {
	Name string `json:"name"`
}

type DepositRequestBody struct {
	Amount float64 `json:"amount"`
}

func HandleRequest(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	bodyString := request.Body

	var body DepositRequestBody
	err := json.Unmarshal([]byte(bodyString), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: fmt.Sprintf("Successfully deposited %.2f", body.Amount),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
