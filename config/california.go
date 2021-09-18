package config

import "strings"

type California struct {
}

func (i California) Color() Color {
	return Color{Blue: 255, Green: 255, Red: 255}
}

func (i California) Prefix() string {
	return "CA"
}

func (i California) Match(pattern string) (SignTemplate, bool) {
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
			Template: "California.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     300,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     230,
					Y:            333,
					Center:       192.5,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "California3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     430,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     230,
					Y:            333,
					Center:       224.5,
				},
			},
		}, true
	}
}
