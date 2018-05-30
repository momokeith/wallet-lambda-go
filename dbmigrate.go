package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func migrateDB(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body := "{\"migrated\":true}"
	return events.APIGatewayProxyResponse{
		Body: body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(migrateDB)
}
