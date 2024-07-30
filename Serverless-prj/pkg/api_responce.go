package handler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponce(status int, body interface{}) (*events.APIGatewayProxyRequest, error) {
	resp := event.APIGatewayProxyRequest{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBod)
	return &resp, nil
}
