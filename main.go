package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang/freetype/truetype"
	"image/png"
	"sign-builder/config"

	"github.com/GeertJohan/go.rice"
	"github.com/fogleman/gg"
)

func HandleShieldQuery ( request events.APIGatewayProxyRequest ) (events.APIGatewayProxyResponse, error) {
	//Make sure fonts are available when script is loaded
	_ = rice.MustFindBox("fonts")

	_ = rice.MustFindBox("templates")

	pattern, ok := request.QueryStringParameters["shield"]
	if !ok {
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: "Count Not Find Parameter Shield" } , nil
	}

	template, err := config.LoadSignTemplate(pattern)

	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
	}

	image, _, err := template.LoadTemplateImage()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
	}

	ctx := gg.NewContextForImage(image)

	for _, textBlock := range template.TextBlocks {
		opts := truetype.Options{}
		opts.Size = textBlock.FontSize
		font, err := config.LoadFont(textBlock.DefaultFont)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
		}
		face := truetype.NewFace(font, &opts)
		textBlock.SetColor(ctx.SetRGBA255)

		ctx.SetFontFace(face)
		fmt.Println(ctx.MeasureString(textBlock.Text))
		w, _ := ctx.MeasureString(textBlock.Text)
		if w > float64(textBlock.MaxWidth) {
			font, err := config.LoadFont(textBlock.OversizeFont)
			if err != nil {
				return events.APIGatewayProxyResponse { StatusCode: 400,Body: err.Error() } , err
			}
			face := truetype.NewFace(font, &opts)
			ctx.SetFontFace(face)
			w, _ = ctx.MeasureString(textBlock.Text)
		}
		newX := textBlock.Position(w)
		ctx.DrawString(textBlock.Text, newX, float64(textBlock.Y))
	}

	var buff bytes.Buffer
	png.Encode(&buff, ctx.Image())

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
