package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/disintegration/imaging"
	"path"
	"sign-builder/core"

	"github.com/GeertJohan/go.rice"
)

var largeHeight = 50
var smallHeight = 20

func init() {
	core.Init()
}

func HandleShieldPostQuery ( request events.APIGatewayProxyRequest ) (events.APIGatewayProxyResponse, error) {
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

	smallImage := imaging.Resize(*img, 0, smallHeight, imaging.Linear)
	smallFileName := path.Join("/tmp", pattern + ".png")
	smallKey := path.Join("Shields", "20x", pattern + ".png")
	err = imaging.Save(smallImage, smallFileName)
	if err != nil {
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
	}

	err = core.UploadS3(smallKey, smallFileName)
	if err != nil {
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
	}

	largeImage := imaging.Resize(*img, 0, largeHeight, imaging.Linear)
	largeFileName := path.Join("/tmp", pattern + ".png")
	largeKey := path.Join("Shields", pattern + ".png")
	err = imaging.Save(largeImage, largeFileName)
	if err != nil {
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
	}
	err = core.UploadS3(largeKey, largeFileName)
	if err != nil {
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}

	return resp, nil
}

func main() {
	lambda.Start(HandleShieldPostQuery)
}
