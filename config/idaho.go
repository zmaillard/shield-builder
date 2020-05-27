package config

import "strings"

type Idaho struct {

}


func (i Idaho) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Idaho) Prefix() string {
	return "ID"
}

func (i Idaho) Match(pattern string) (SignTemplate, bool) {
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



	if twoDigitShield {
		var rightAlign = 580.0
		return SignTemplate{
			Template: "Idaho.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     300,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     375,
					Y:            280,
					Right:        &rightAlign,
				},
			},
		}, true
	} else {
		var rightAlign = 745.0
		return SignTemplate{
			Template: "Idaho3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     600,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     370,
					Y:            280,
					Right:        &rightAlign,
				},
			},
		}, true
	}

}
