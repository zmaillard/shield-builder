package config

import "strings"

type BusinessLoop struct {
}

func (i BusinessLoop) Color() Color {
	return Color{Blue: 255, Green: 255, Red: 255}
}

func (i BusinessLoop) Prefix() string {
	return "BL"
}

func (i BusinessLoop) Match(pattern string) (SignTemplate, bool) {
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
			Template: "BusinessLoop.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     380,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     310,
					Y:            430,
					Center:       300.0,
				},
			},
		}, true
	}

	return SignTemplate{}, false
}

