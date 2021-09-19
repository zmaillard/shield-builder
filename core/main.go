package core

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	gofont "golang.org/x/image/font"
	"image"
	"sign-builder/config"
)

func Build(pattern string) (*image.Image, error) {
	template, err := config.LoadSignTemplate(pattern)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	image, _, err := template.LoadTemplateImage()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ctx := gg.NewContextForImage(image)

	for _, textBlock := range template.TextBlocks {
		opts := truetype.Options{}
		opts.Size = textBlock.FontSize
		font, err := config.LoadFont(textBlock.DefaultFont)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		face := truetype.NewFace(font, &opts)
		textBlock.SetColor(ctx.SetRGBA255)

		ctx.SetFontFace(face)
		w, _ := ctx.MeasureString(textBlock.Text)
		var face2 gofont.Face
		if textBlock.ShrinkLastCharacterBy != nil {
			opts := truetype.Options{}

			opts.Size = *textBlock.ShrinkLastCharacterBy
			font, err := config.LoadFont(textBlock.DefaultFont)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			face2 = truetype.NewFace(font, &opts)
			w1, _ := ctx.MeasureString(textBlock.Text[:len(textBlock.Text)-1])
			ctx.SetFontFace(face2)
			w2, _ := ctx.MeasureString(textBlock.Text[len(textBlock.Text)-1:])
			w = w1 + w2
			ctx.SetFontFace(face)
		}

		if w > float64(textBlock.MaxWidth) {
			font, err := config.LoadFont(textBlock.OversizeFont)
			if err != nil {
				return nil, err
			}
			face := truetype.NewFace(font, &opts)
			ctx.SetFontFace(face)
			w, _ = ctx.MeasureString(textBlock.Text)
			if textBlock.ShrinkLastCharacterBy != nil {
				w1, _ := ctx.MeasureString(textBlock.Text[:len(textBlock.Text)-1])
				ctx.SetFontFace(face2)
				w2, _ := ctx.MeasureString(textBlock.Text[len(textBlock.Text)-1:])
				w = w1 + w2
				ctx.SetFontFace(face)
			}
		}
		newX := textBlock.Position(w)
		if textBlock.ShrinkLastCharacterBy == nil {
			ctx.DrawString(textBlock.Text, newX, float64(textBlock.Y))
		} else {
			w1, _ := ctx.MeasureString(textBlock.Text[:len(textBlock.Text)-1])
			ctx.DrawString(textBlock.Text[:len(textBlock.Text)-1], newX, float64(textBlock.Y))
			ctx.SetFontFace(face2)
			newX = newX + w1
			ctx.DrawString(textBlock.Text[len(textBlock.Text)-1:], newX, float64(textBlock.Y))
		}

	}

	img := ctx.Image()
	return &img, nil
}
