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
	isGuide := false

	if len(suffixTokens) == 2 {
		if strings.ToUpper(suffixTokens[1]) == "GUIDE" {
			isGuide = true
		}
	}

	// Number
	twoDigitShield := true
	if len(suffixTokens[0]) > 2 {
		twoDigitShield = false
	}
	if isGuide && twoDigitShield {
		// Two digit Guide shield
		return SignTemplate{
			Template: "WyomingGuide.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     380,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					FontSize:     275,
					Y:            330,
					Center:       197,
				},
			},
		}, true
	} else if isGuide && !twoDigitShield {
		// Three digit Guide shield
		return SignTemplate{
			Template: "WyomingGuide3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     385,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					FontSize:     273,
					Y:            340,
					Center:       256,
				},
			},
		}, true
	} else if twoDigitShield {
		// Two digit shield without state
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