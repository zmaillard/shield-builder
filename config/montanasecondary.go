package config

import "strings"


type MontanaSecondary struct {

}

func (i MontanaSecondary) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i MontanaSecondary) Prefix() string {
	return "MTS"
}

func (i MontanaSecondary) Match(pattern string) (SignTemplate, bool) {
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

	if !isGuide {
		return SignTemplate{
			Template: "MontanaSecondary.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     350,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     250,
					Y:            355,
					Center:       300,
				},
			},
		}, true
	} else   {
		return SignTemplate{
			Template: "MontanaSecondaryGuide.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     500,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         number,
					Color:        i.Color(),
					FontSize:     250,
					Y:            345,
					Center:       300,
				},
			},
		}, true
	}
}