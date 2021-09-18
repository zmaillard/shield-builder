package config

import "strings"

type Montana struct {
}

func (i Montana) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Montana) Prefix() string {
	return "MT"
}

func (i Montana) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	// Number
	twoDigitShield := true

	suffixTokens := strings.Split(tokens[1], "_")
	isGuide := false
	number := tokens[1]
	if len(suffixTokens) == 2 && strings.ToUpper(suffixTokens[1]) == "GUIDE" {
		isGuide = true
		number = suffixTokens[0]
	}

	if len(number) > 2 {
		twoDigitShield = false
	}

	if twoDigitShield && !isGuide {
		return SignTemplate{
			Template: "Montana.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     550,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     375,
					Y:            470,
					Center:       300,
				},
			},
		}, true
	} else if !twoDigitShield && !isGuide {
		return SignTemplate{
			Template: "Montana3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     700,
					DefaultFont:  D,
					OversizeFont: B,
					Text:         number,
					Color:        i.Color(),
					FontSize:     375,
					Y:            470,
					Center:       375,
				},
			},
		}, true
	} else if twoDigitShield && isGuide {
		return SignTemplate{
			Template: "MontanaGuide.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     580,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     425,
					Y:            450,
					Center:       300,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "MontanaGuide3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     740,
					DefaultFont:  D,
					OversizeFont: B,
					Text:         number,
					Color:        i.Color(),
					FontSize:     400,
					Y:            430,
					Center:       375,
				},
			},
		}, true
	}
}
