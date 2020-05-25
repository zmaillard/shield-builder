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
	if len(tokens[1]) > 2 {
		twoDigitShield = false
	}

	// Two digit shield without state
	if twoDigitShield {
		return SignTemplate{
			Template: "Utah.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     240,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     192.5,
					Y:            250,
					Center:       195.35,
				},
			},
		}, true
	} else  {
		return SignTemplate{
			Template: "Utah.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     240,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     160,
					Y:            250,
					Center:       195.35,
				},
			},
		}, true
	}
}