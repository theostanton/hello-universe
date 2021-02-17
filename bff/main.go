package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
)

func createResponse(success bool, message string, version string) (events.APIGatewayProxyResponse, error) {

	type Body struct {
		Success bool   `json:"success"`
		Message string `json:"message,omitempty"`
		Version string `json:"version,omitempty"`
	}

	body := Body{
		Success: success,
		Message: message,
		Version: version,
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json; charset=utf-8"},
		Body:       string(bodyBytes),
	}, nil
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("handleRequest()")

	apiVersion, exists := os.LookupEnv("API_VERSION")

	if !exists {
		return createResponse(false, "No API_VERSION environment variable", "")
	}

	return createResponse(true, "Hello World", apiVersion)
}

func main() {
	lambda.Start(handleRequest)
}
