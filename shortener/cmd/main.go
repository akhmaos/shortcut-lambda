package main

import (
	"github.com/akhmaos/shortcut-lambda/shortener/internal"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(internal.Handler)
}
