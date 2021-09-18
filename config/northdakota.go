package config

import (
	"strings"
)

func getNorthDakotaShieldWidth(number string) ShieldWidth {
	switch len(number) {
	case 1:
		return TwoDigit
	case 2:
		return TwoDigit
	case 3:
		return ThreeDigit
	case 4:
		return FourDigit
	default:
		return TwoDigit
	}

	return TwoDigit
}

type NorthDakota struct {
}

func (i NorthDakota) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i NorthDakota) Prefix() string {
	return "ND"
}

func (i NorthDakota) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	shieldWidth := getNorthDakotaShieldWidth(tokens[1])

	// Two digit shield without state
	if shieldWidth == TwoDigit {
		return SignTemplate{
			Template: "NorthDakota.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     580,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     448,
					Y:            510,
					Center:       275,
				},
			},
		}, true
	} else if shieldWidth == ThreeDigit {
		return SignTemplate{
			Template: "NorthDakota3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     805,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     448,
					Y:            510,
					Center:       355,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "NorthDakota3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     805,
					DefaultFont:  B,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     448,
					Y:            510,
					Center:       355,
				},
			},
		}, true
	}
}
