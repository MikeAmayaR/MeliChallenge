package main

import (
	handler "github.com/MikeAmayaR/MeliChallenge.git/internal/api"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
