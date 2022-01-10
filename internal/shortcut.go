package internal

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func initDB() (*DBSession, error) {
	return &DBSession{}, nil
}

func ApiGatewayHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{Body: "s.hach.dev/test", StatusCode: 200}, nil
}
