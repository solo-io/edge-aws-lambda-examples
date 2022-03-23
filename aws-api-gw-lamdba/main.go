package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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

	responseBodyString, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error Marshalling Request Body.")
		ApiResponse = events.APIGatewayProxyResponse{Body: err.Error(), Headers: headers, IsBase64Encoded: false, MultiValueHeaders: returnHeaderExample, StatusCode: 500}
		return ApiResponse, nil
	}
	statusCode := 200
	responseCode := queryStringParameters["responseCode"]
	if responseCode != "" {
		rc, err := strconv.Atoi(responseCode)
		if err != nil {
			statusCode = 500
			fmt.Println("Error Converting responseCode to int.  Setting statusCode to 500.")
		} else {
			statusCode = rc
		}
	}
	ApiResponse = events.APIGatewayProxyResponse{Body: string(responseBodyString), Headers: headers, IsBase64Encoded: false, MultiValueHeaders: returnHeaderExample, StatusCode: statusCode}
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
