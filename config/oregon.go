package config

import "strings"

type Oregon struct {
}

func (i Oregon) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Oregon) Prefix() string {
	return "OR"
}

func (i Oregon) Match(pattern string) (SignTemplate, bool) {
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
			Template: "Oregon.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     530,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     425,
					Y:            405,
					Center:       300,
				},
			},
		}, true
	} else  {
		return SignTemplate{
			Template: "Oregon3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     660,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     425,
					Y:            405,
					Center:       375,
				},
			},
		}, true
	}
}
