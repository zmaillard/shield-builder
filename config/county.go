package config

import "strings"

type County struct {
}

func (i County) Color() Color {
	return Color{Blue: 23, Green: 209, Red: 247}
}

func (i County) Prefix() string {
	return "CH"
}

func (i County) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	return SignTemplate{
		Template: "CountyHighway.png",
		TextBlocks: []TextBlock{
			{
				MaxWidth:     525,
				DefaultFont:  D,
				OversizeFont: C,
				Text:         tokens[1],
				Color:        i.Color(),
				FontSize:     415,
				Y:            575,
				Center:       400,
			},
		},
	}, true
}
