package config

import "strings"

type Washington struct {
}

func (i Washington) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Washington) Prefix() string {
	return "WA"
}

func (i Washington) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	// Number
	twoDigitShield := true
	if len(tokens[1]) > 2 {
		twoDigitShield = false
	}

	// Two digit shield without state
	if twoDigitShield {
		return SignTemplate{
			Template: "Washington.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     490,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     400,
					Y:            390,
					Center:       300,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "Washington.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     450,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     240,
					Y:            348,
					Center:       300,
				},
			},
		}, true
	}
}
