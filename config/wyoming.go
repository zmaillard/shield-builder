package config

import "strings"

type Wyoming struct {

}

func (i Wyoming) Color() Color {
	return Color{Blue: 13, Green: 13, Red: 38}
}

func (i Wyoming) Prefix() string {
	return "WY"
}

func (i Wyoming) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}


	// Check Suffix
	suffixTokens := strings.Split(tokens[1], "_")
	/*
	isGuide = false



	if len(suffixTokens) == 2 {
		if strings.ToUpper(suffixTokens[1]) == "GUIDE" {
			isGuide = true
		}
	}

	 */

	// Number
	twoDigitShield := true
	if len(suffixTokens[0]) > 2 {
		twoDigitShield = false
	}

	// Two digit shield without state
	if twoDigitShield {
		return SignTemplate{
			Template: "Wyoming.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     340,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					FontSize:     210,
					Y:            283,
					Center:       195.35,
				},
			},
		}, true
	} else  {
		return SignTemplate{
			Template: "Wyoming.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     340,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					FontSize:     210,
					Y:            283,
					Center:       195.35,
				},
			},
		}, true
	}

}