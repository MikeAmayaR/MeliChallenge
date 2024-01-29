package handler

import (
	"context"
	"encoding/json"

	"github.com/MikeAmayaR/MeliChallenge.git/internal/model"
	"github.com/MikeAmayaR/MeliChallenge.git/internal/service"
	"github.com/aws/aws-lambda-go/events"
)

func HandleTopSecret(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var req model.TopSecretRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return ApiResponse(400, "Invalid payload"), nil
	}

	distances := make([]float32, 0)
	messages := make([][]string, 0)

	for _, satellite := range req.Satellites {
		distances = append(distances, satellite.Distance)
		messages = append(messages, satellite.Message)
	}

	x, y, err := service.GetLocation(distances...)
	msg := service.GetMessage(messages...)

	if x == 0 && y == 0 || msg == "" {
		return ApiResponse(404, "Not found"), nil
	}

	response := model.TopSecretResponse{
		Position: model.Position{
			X: x,
			Y: y,
		},
		Message: msg,
	}

	return ApiResponse(200, response), nil
}

func ApiResponse(status int, body interface{}) *events.APIGatewayProxyResponse {
	resp := &events.APIGatewayProxyResponse{
		StatusCode: status,
	}

	var stringBody string

	switch b := body.(type) {
	case string:
		stringBody = b
	default:
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       err.Error(),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			}
		}
		stringBody = string(bodyBytes)
	}

	resp.Body = stringBody

	if status >= 200 && status < 300 {
		resp.Headers = map[string]string{
			"Content-Type": "application/json",
		}
	}

	return resp
}
