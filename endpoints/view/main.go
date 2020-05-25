package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"image/png"
	"sign-builder/core"

	"github.com/GeertJohan/go.rice"
)

func HandleShieldQuery ( request events.APIGatewayProxyRequest ) (events.APIGatewayProxyResponse, error) {
	//Make sure fonts are available when script is loaded
	_ = rice.MustFindBox("fonts")

	_ = rice.MustFindBox("templates")

	pattern, ok := request.QueryStringParameters["shield"]
	if !ok {
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: "Count Not Find Parameter Shield" } , nil
	}

	img, err := core.Build(pattern)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
	}

	var buff bytes.Buffer
	png.Encode(&buff, *img)

	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	headers := map[string]string  {
		"Content-type": "image/png",
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: headers,
		Body: encodedString,
		IsBase64Encoded: true,
	}

	return resp, nil
}

func main() {
	lambda.Start(HandleShieldQuery)
}
