package main

import (
	"fmt"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type bodyResponse struct {
	Resource              string            `json:"resource"` // The resource path defined in API Gateway
	Path                  string            `json:"path"`     // The url path for the caller
	HTTPMethod            string            `json:"httpMethod"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	PathParameters        map[string]string `json:"pathParameters"`
	StageVariables        map[string]string `json:"stageVariables"`
	Body                  string            `json:"body"`
	IsBase64Encoded       bool              `json:"isBase64Encoded,omitempty"`
}

/*
type APIGatewayProxyRequest struct {
    Resource              string                        `json:"resource"` // The resource path defined in API Gateway
    Path                  string                        `json:"path"`     // The url path for the caller
    HTTPMethod            string                        `json:"httpMethod"`
    Headers               map[string]string             `json:"headers"`
    QueryStringParameters map[string]string             `json:"queryStringParameters"`
    PathParameters        map[string]string             `json:"pathParameters"`
    StageVariables        map[string]string             `json:"stageVariables"`
    RequestContext        APIGatewayProxyRequestContext `json:"requestContext"`
    Body                  string                        `json:"body"`
    IsBase64Encoded       bool                          `json:"isBase64Encoded,omitempty"`
}
type APIGatewayProxyResponse struct {
    StatusCode      int               `json:"statusCode"`
    Headers         map[string]string `json:"headers"`
    Body            string            `json:"body"`
    IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}
*/

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}
	returnHeaderExample := map[string][]string{"x-multivalue-header": []string{"multivalue-header-1", "multivalue-header-2", "multivalue-header-3"}}
	resource := request.Resource
	path := request.Path
	httpMethod := request.HTTPMethod
	headers := request.Headers
	queryStringParameters := request.QueryStringParameters
	pathParameters := request.PathParameters
	stageVariables := request.StageVariables
	body := request.Body
	isBase64Encoded := request.IsBase64Encoded

	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(body))
	fmt.Printf("Resource = %s.\n", resource)
	fmt.Printf("Path = %s.\n", path)
	fmt.Printf("HttpMethod = %s.\n", httpMethod)
	fmt.Printf("Body = %s.\n", body)
	fmt.Printf("IsBase64Encoded = %t.\n", isBase64Encoded)
	fmt.Println("Headers:")
	for headerKey, headerValue := range headers {
		fmt.Printf("    %s: %s\n", headerKey, headerValue)
	}
	fmt.Println("QueryStringParameters:")
	for qsKey, qsValue := range queryStringParameters {
		fmt.Printf("    %s: %s\n", qsKey, qsValue)
	}
	fmt.Println("PathParameters:")
	for pathKey, pathValue := range pathParameters {
		fmt.Printf("    %s: %s\n", pathKey, pathValue)
	}
	fmt.Println("StageVariables:")
	for stageKey, stageValue := range stageVariables {
		fmt.Printf("    %s: %s\n", stageKey, stageValue)
	}

	responseBody := &bodyResponse{
		Resource:              resource,
		Path:                  path,
		HTTPMethod:            httpMethod,
		Headers:               headers,
		QueryStringParameters: queryStringParameters,
		PathParameters:        pathParameters,
		StageVariables:        stageVariables,
		Body:                  body,
		IsBase64Encoded:       isBase64Encoded,
	}

	responseBodyString, err := json.Marshal(responseBody)
	if err != nil {
		fmt.Println("Error Marshalling Response Body.")
	}

	ApiResponse = events.APIGatewayProxyResponse{Body: string(responseBodyString), Headers: headers, IsBase64Encoded: false, MultiValueHeaders: returnHeaderExample, StatusCode: 200}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
