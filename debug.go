package main

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image/png"
	"os"
	"sign-builder/config"
)

func main(){
	//Make sure fonts are available when script is loaded
	_ = rice.MustFindBox("fonts")

	_ = rice.MustFindBox("templates")

	pattern := os.Args[1]

	template, err := config.LoadSignTemplate(pattern)

	if err != nil {
		fmt.Println(err)
		return
	}

	image, _, err := template.LoadTemplateImage()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := gg.NewContextForImage(image)

	for _, textBlock := range template.TextBlocks {
		opts := truetype.Options{}
		opts.Size = textBlock.FontSize
		font, err := config.LoadFont(textBlock.DefaultFont)
		if err != nil {
			fmt.Println(err)
			return
		}
		face := truetype.NewFace(font, &opts)
		textBlock.SetColor(ctx.SetRGBA255)

		ctx.SetFontFace(face)
		fmt.Println(ctx.MeasureString(textBlock.Text))
		w, _ := ctx.MeasureString(textBlock.Text)
		if w > float64(textBlock.MaxWidth) {
			font, err := config.LoadFont(textBlock.OversizeFont)
			if err != nil {
				return
			}
			face := truetype.NewFace(font, &opts)
			ctx.SetFontFace(face)
			w, _ = ctx.MeasureString(textBlock.Text)
		}
		newX := textBlock.Position(w)
		ctx.DrawString(textBlock.Text, newX, float64(textBlock.Y))
	}


	f, err := os.Create("/tmp/" + pattern + ".png")
	if err != nil {
		return
	}
	defer f.Close()

	png.Encode(f, ctx.Image())

}
