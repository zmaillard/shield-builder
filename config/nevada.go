package config

import "strings"

type Nevada struct {

}


func (i Nevada) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Nevada) Prefix() string {
	return "NV"
}

func (i Nevada) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	suffixTokens := strings.Split(tokens[1], "_")
	isGuide := false
	number := tokens[1]
	if len(suffixTokens) == 2 && strings.ToUpper(suffixTokens[1]) == "GUIDE" {
		isGuide = true
		number = suffixTokens[0]
	}

	if isGuide {
		return SignTemplate{
			Template: "Nevada_Guide.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     775,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         number,
					Color:        i.Color(),
					FontSize:     440,
					Y:            380,
					Center:       470,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "Nevada.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     380,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     250,
					Y:            230,
					Center:       300,
				},
			},
		}, true
	}


}