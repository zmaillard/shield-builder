package config

import "strings"

type Minnesota struct {
}

func (i Minnesota) Color() Color {
	return Color{Blue: 255, Green: 255, Red: 255}
}

func (i Minnesota) Prefix() string {
	return "MN"
}

func (i Minnesota) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	// Number
	twoDigitShield := true

	number := tokens[1]

	if len(number) > 2 {
		twoDigitShield = false
	}

	// Two digit shield without state
	if twoDigitShield {
		return SignTemplate{
			Template: "Minnesota.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     700,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     460,
					Y:            610,
					Center:       375,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "Minnesota3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     920,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         number,
					Color:        i.Color(),
					FontSize:     460,
					Y:            610,
					Center:       469,
				},
			},
		}, true
	}
}

type MinnesotaCounty struct {
}

func (i MinnesotaCounty) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i MinnesotaCounty) Prefix() string {
	return "MNCH"
}

func (i MinnesotaCounty) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}
	number := tokens[1]

	return SignTemplate{
		Template: "MinnesotaCounty.png",
		TextBlocks: []TextBlock{
			{
				MaxWidth:     775,
				DefaultFont:  D,
				OversizeFont: C,
				Text:         number,
				Color:        i.Color(),
				FontSize:     490,
				Y:            680,
				Center:       400,
			},
		},
	}, true

}
