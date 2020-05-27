package core

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image"
	"sign-builder/config"
)


func Build(pattern string) (*image.Image, error) {
	template, err := config.LoadSignTemplate(pattern)

	if err != nil {
		fmt.Println(err)
		return nil , err
	}

	image, _, err := template.LoadTemplateImage()
	if err != nil {
		fmt.Println(err)
		return nil , err
	}

	ctx := gg.NewContextForImage(image)

	for _, textBlock := range template.TextBlocks {
		opts := truetype.Options{}
		opts.Size = textBlock.FontSize
		font, err := config.LoadFont(textBlock.DefaultFont)
		if err != nil {
			fmt.Println(err)
			return nil , err
		}
		face := truetype.NewFace(font, &opts)
		textBlock.SetColor(ctx.SetRGBA255)

		ctx.SetFontFace(face)
		fmt.Println(ctx.MeasureString(textBlock.Text))
		w, _ := ctx.MeasureString(textBlock.Text)
		if w > float64(textBlock.MaxWidth) {
			font, err := config.LoadFont(textBlock.OversizeFont)
			if err != nil {
				return  nil, err
			}
			face := truetype.NewFace(font, &opts)
			ctx.SetFontFace(face)
			w, _ = ctx.MeasureString(textBlock.Text)
		}
		newX := textBlock.Position(w)
		ctx.DrawString(textBlock.Text, newX, float64(textBlock.Y))
	}

	img := ctx.Image()
	return &img, nil
}
