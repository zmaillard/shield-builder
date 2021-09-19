package config

import "strings"

type NewMexico struct {
}

func (i NewMexico) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i NewMexico) Prefix() string {
	return "NM"
}

func (i NewMexico) Match(pattern string) (SignTemplate, bool) {
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
			Template: "New_Mexico.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     360,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     285,
					Y:            390,
					Center:       300,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "New_Mexico.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     360,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     225,
					Y:            390,
					Center:       300,
				},
			},
		}, true
	}
}
