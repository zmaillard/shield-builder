package config

import "strings"

type US struct {
}

func (i US) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i US) Prefix() string {
	return "US"
}

func (i US) CaliforniaSuffix() string {
	return "CA"
}

func (i US) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	var isCalifornia = false
	// Check Suffix
	suffixTokens := strings.Split(tokens[1], "_")
	if len(suffixTokens) == 2 {
		if suffixTokens[1] == i.CaliforniaSuffix() {
			isCalifornia = true
		}

	}

	// Number
	twoDigitShield := true
	if len(suffixTokens[0]) > 2 {
		twoDigitShield = false
	}

	// Two digit shield without state
	if twoDigitShield && !isCalifornia {
		return SignTemplate{
			Template: "US.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     550,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					FontSize:     450,
					Y:            440,
					Center:       300,
				},
			},
		}, true
	} else if twoDigitShield && isCalifornia {
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
			},
		}, true
	} else if !twoDigitShield && !isCalifornia {
		return SignTemplate{
			Template: "US3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     660,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					Center:       375,
					FontSize:     425,
					Y:            425,
				},
			},
		}, true
	} else if !twoDigitShield && isCalifornia {
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
			},
		}, true
	}

	return SignTemplate{}, false
}

