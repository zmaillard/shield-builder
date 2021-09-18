package config

import "strings"

type Utah struct {
}

func (i Utah) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Utah) Prefix() string {
	return "UT"
}

func (i Utah) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	// Number
	twoDigitShield := true

	suffixTokens := strings.Split(tokens[1], "_")
	isGuide := false
	number := tokens[1]
	if len(suffixTokens) == 2 && strings.ToUpper(suffixTokens[1]) == "GUIDE" {
		isGuide = true
		number = suffixTokens[0]
	}

	if len(number) > 2 {
		twoDigitShield = false
	}

	// Two digit shield without state
	if twoDigitShield && !isGuide {
		return SignTemplate{
			Template: "Utah.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     240,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     192.5,
					Y:            250,
					Center:       195.35,
				},
			},
		}, true
	} else if !twoDigitShield && !isGuide {
		return SignTemplate{
			Template: "Utah.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     240,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         number,
					Color:        i.Color(),
					FontSize:     160,
					Y:            250,
					Center:       195.35,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "UtahGuide.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     450,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     300,
					Y:            385,
					Center:       335,
				},
			},
		}, true
	}
}
