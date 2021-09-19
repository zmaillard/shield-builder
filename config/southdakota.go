package config

import (
	"strings"
)

type ShieldWidth int

const (
	TwoDigit ShieldWidth = iota
	ThreeDigit
	FourDigit
)

func getShieldWidth(number string) ShieldWidth {
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

type SouthDakota struct {
}

func (i SouthDakota) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i SouthDakota) Prefix() string {
	return "SD"
}

func (i SouthDakota) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	shieldWidth := getShieldWidth(tokens[1])

	// Two digit shield without state
	if shieldWidth == TwoDigit {
		return SignTemplate{
			Template: "SouthDakota.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     550,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     450,
					Y:            440,
					Center:       300,
				},
			},
		}, true
	} else if shieldWidth == ThreeDigit {
		return SignTemplate{
			Template: "SouthDakota3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     620,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     448,
					Y:            440,
					Center:       375,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "SouthDakota3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     620,
					DefaultFont:  B,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     448,
					Y:            440,
					Center:       375,
				},
			},
		}, true
	}
}
