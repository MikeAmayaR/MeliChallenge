package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch {
	case request.HTTPMethod == http.MethodPost && request.Resource == "/topsecret":
		return HandleTopSecret(ctx, request)

	case request.HTTPMethod == http.MethodPost && strings.HasPrefix(request.Resource, "/topsecret_split/"):
		return PostTopSecretSplitHandler(ctx, request)

	case request.HTTPMethod == http.MethodGet && request.Resource == "/topsecret_split":
		return GetTopSecretSplitHandler(ctx, request)

	default:
		return ApiResponse(http.StatusNotFound, "Resource not found"), nil
	}
}
