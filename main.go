package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
     _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/momokeith/wallet-lambda-go/wallet"
)

func handleRequest(context context.Context,
	request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	wallets := wallet.List()

	return events.APIGatewayProxyResponse{
		Body: wallets.Json(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}