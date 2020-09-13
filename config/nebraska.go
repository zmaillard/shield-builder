package config

import "strings"

type Nebraska struct {

}


func (i Nebraska) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Nebraska) Prefix() string {
	return "NE"
}

func (i Nebraska) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	suffixTokens := strings.Split(tokens[1], "_")
	isLink := false
	isSpur := false
	linkSpurNumber := ""
	if len(suffixTokens) == 2 && strings.ToUpper(suffixTokens[1]) == "LINK" {
		linkSpurNumber = suffixTokens[0]
		isLink = true
	} else if len(suffixTokens) == 2 && strings.ToUpper(suffixTokens[1]) == "SPUR" {
		linkSpurNumber = suffixTokens[0]
		isSpur = true
	}

	// Number
	twoDigitShield := true
	if len(tokens[1]) > 2 {
		twoDigitShield = false
	}

	isNormal := !isLink && !isSpur

	shrinkBy :=210.0

	if twoDigitShield && isNormal {
		return SignTemplate{
			Template: "Nebraska.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     475,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     285,
					Y:            300,
					Center:       250,
				},
			},
		}, true
	} else if !twoDigitShield && isNormal {
		return SignTemplate{
			Template: "Nebraska3di.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     475,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     215,
					Y:            235,
					Center:       250,
				},
			},
		}, true
	} else if isSpur {
		return SignTemplate{
			Template: "NebraskaSpur.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     400,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         linkSpurNumber,
					Color:        i.Color(),
					FontSize:     280,
					Y:            250,
					Center:       250,
					ShrinkLastCharacterBy: &shrinkBy,
				},
			},
		}, true
	}else  { //link
		return SignTemplate{
			Template: "NebraskaLink.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     400,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         linkSpurNumber,
					Color:        i.Color(),
					FontSize:     280,
					Y:            250,
					Center:       250,
					ShrinkLastCharacterBy: &shrinkBy,
				},
			},
		}, true
	}

}
