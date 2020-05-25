package config

import "strings"

type Interstate struct {
}

func (i Interstate) Color() Color {
	return Color{Blue: 255, Green: 255, Red: 255}
}

func (i Interstate) Prefix() string {
	return "I"
}

func (i Interstate) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	var stateName = ""

	// Check Suffix
	suffixTokens := strings.Split(tokens[1], "_")
	if len(suffixTokens) == 2 {
		if state, ok := stateSuffix[suffixTokens[1]]; ok {
			stateName = strings.ToUpper(state)
		} else {
			return SignTemplate{}, false
		}
	}

	// Number
	twoDigitShield := true
	if len(suffixTokens[0]) > 2 {
		twoDigitShield = false
	}

	// Two digit shield without state
	if twoDigitShield && len(stateName) == 0 {
		return SignTemplate{
			Template: "Interstate.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     295,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					FontSize:     250,
					Y:            300,
					Center:       192.5,
				},
			},
		}, true
	} else if twoDigitShield && len(stateName) > 0 {
		// Two digit shield with state
		return SignTemplate{
			Template: "Interstate.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     295,
					FontSize:     185,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					Y:            300,
					Center:       192.5,
				},
				{
					MaxWidth:     295,
					FontSize:     33,
					DefaultFont:  E,
					OversizeFont: D,
					Text:         stateName,
					Color:        i.Color(),
					Y:            150,
					Center:       192.5,
				},
			},
		}, true
	} else if !twoDigitShield && len(stateName) == 0 {
		return SignTemplate{
			Template: "Interstate3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     490,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					Center:       375.5,
					FontSize:     400,
					Y:            470,
				},
			},
		}, true
	} else if !twoDigitShield && len(stateName) > 0 {
		return SignTemplate{
			Template: "Interstate3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     490,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					Center:       375.5,
					FontSize:     343,
					Y:            470,
				},
				{
					MaxWidth:     490,
					DefaultFont:  E,
					OversizeFont: D,
					Text:         stateName,
					Color:        i.Color(),
					Center:       375.5,
					FontSize:     62,
					Y:            200,
				},
			},
		}, true
	}

	return SignTemplate{}, false
}
