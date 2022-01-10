package main

import (
	"github.com/akhmaos/shortcut-lambda/internal"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(internal.ApiGatewayHandler)
}
