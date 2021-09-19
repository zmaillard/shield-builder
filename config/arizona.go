package config

import "strings"

type Arizona struct {
}

func (i Arizona) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Arizona) LoopSuffix() string {
	return "L"
}

func (i Arizona) Prefix() string {
	return "AZ"
}

func (i Arizona) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	hasLoop := false
	if strings.Index(tokens[1], i.LoopSuffix()) >= 0 {
		hasLoop = true
		tokens[1] = strings.ReplaceAll(tokens[1], i.LoopSuffix(), "")
	}

	// Number
	twoDigitShield := true
	if len(tokens[1]) > 2 {
		twoDigitShield = false
	}

	if twoDigitShield {
		return SignTemplate{
			Template: "Arizona.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     540,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     440,
					Y:            475,
					Center:       300,
				},
			},
		}, true
	} else if hasLoop {
		return SignTemplate{
			Template: "ArizonaLoop.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     690,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     440,
					Y:            475,
					Center:       375,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "Arizona3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     690,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     440,
					Y:            475,
					Center:       375,
				},
			},
		}, true
	}

}
