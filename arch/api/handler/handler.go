package main

import (
	"context"
	"easyparking/app/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs

	ginLambda = ginadapter.New(handler.Server)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return ginLambda.ProxyWithContext(ctx, req)

}

func main() {
	lambda.Start(Handler)
}
