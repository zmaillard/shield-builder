package config

import "strings"

type Idaho struct {

}


func (i Idaho) Color() Color {
	return Color{Blue: 255, Green: 255, Red: 255}
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


	var rightAlign = 225.0
	if twoDigitShield {
		return SignTemplate{
			Template: "Idaho.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     150,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     155,
					Y:            125,
					Right:        &rightAlign,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "Idaho.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     200,
					DefaultFont:  B,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     130,
					Y:            110,
					Right:        &rightAlign,
				},
			},
		}, true
	}

}
