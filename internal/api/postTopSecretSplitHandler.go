package handler

import (
	"context"
	"encoding/json"

	"github.com/MikeAmayaR/MeliChallenge.git/internal/model"
	"github.com/MikeAmayaR/MeliChallenge.git/internal/storage"
	"github.com/aws/aws-lambda-go/events"
)

func PostTopSecretSplitHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	satelliteName, ok := request.PathParameters["satellite_name"]
	if !ok || satelliteName == "" {
		return ApiResponse(400, "Satellite name is required"), nil
	}

	var satelliteData model.Satellite
	err := json.Unmarshal([]byte(request.Body), &satelliteData)
	if err != nil {
		return ApiResponse(400, "Invalid payload"), nil
	}
	satelliteData.Name = satelliteName

	memStorage := storage.GetMemoryStorage()
	memStorage.SaveSatelliteData(satelliteData)

	return ApiResponse(200, "Data saved successfully"), nil
}
