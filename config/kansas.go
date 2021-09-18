package config

import "strings"

type Kansas struct {
}

func (i Kansas) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Kansas) Prefix() string {
	return "KS"
}

func (i Kansas) Match(pattern string) (SignTemplate, bool) {
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
			Template: "Kansas.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     500,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     465,
					Y:            455,
					Center:       300,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "Kansas3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     650,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     465,
					Y:            455,
					Center:       375,
				},
			},
		}, true
	}
}
