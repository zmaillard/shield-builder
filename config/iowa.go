package config

import "strings"

type Iowa struct {
}

func (i Iowa) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Iowa) Prefix() string {
	return "IA"
}

func (i Iowa) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	return SignTemplate{
		Template: "Circle.png",
		TextBlocks: []TextBlock{
			{
				MaxWidth:     680,
				DefaultFont:  D,
				OversizeFont: C,
				Text:         tokens[1],
				Color:        i.Color(),
				FontSize:     450,
				Y:            500,
				Center:       350,
			},
		},
	}, true
}
