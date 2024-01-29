package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MikeAmayaR/MeliChallenge.git/internal/service"
	"github.com/MikeAmayaR/MeliChallenge.git/internal/storage"
	"github.com/aws/aws-lambda-go/events"
)

type successResponse struct {
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func GetTopSecretSplitHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	s := storage.GetMemoryStorage()
	if !s.HasSufficientData() {
		return makeErrorResponse("No hay suficiente información de satélites", http.StatusNotFound)
	}

	allSatelliteData := s.RetrieveAllSatelliteData()

	distancesF32 := make([]float32, len(allSatelliteData))
	messages := [][]string{}
	for i, satData := range allSatelliteData {
		distancesF32[i] = float32(satData.Distance)
		messages = append(messages, satData.Message)
	}

	x, y, err := service.GetLocation(distancesF32...)
	if err != nil {
		return makeErrorResponse("Error calculando la posición", http.StatusInternalServerError)
	}

	message := service.GetMessage(messages...)
	if message == "" {
		return makeErrorResponse("No se pudo determinar el mensaje", http.StatusInternalServerError)
	}

	response := successResponse{
		Position: struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		}{
			X: float64(x),
			Y: float64(y),
		},
		Message: message,
	}

	respBytes, err := json.Marshal(response)
	if err != nil {
		return makeErrorResponse("Error serializando la respuesta", http.StatusInternalServerError)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respBytes),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func makeErrorResponse(message string, statusCode int) (*events.APIGatewayProxyResponse, error) {
	errorResponse, _ := json.Marshal(ErrorResponse{Error: message})
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(errorResponse),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
