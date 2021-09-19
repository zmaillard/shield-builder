package config

import (
	"strconv"
	"strings"
)

type Alberta struct {
}

func (i Alberta) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Alberta) CrowsnestColor() Color {
	return Color{Blue: 255, Green: 255, Red: 255}
}

func (i Alberta) IsSecondary(route int) bool {
	return route >= 500
}

func (i Alberta) IsCrowsNest(route int) bool {
	return route == 3
}

func (i Alberta) Prefix() string {
	return "AB"
}

func (i Alberta) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}

	numberIndex := 0
	route := tokens[1]
	for i := len(route) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(route[0 : i+1])

		if err == nil {
			numberIndex = i
			break
		}
	}

	routeNumber, _ := strconv.Atoi(route[0 : numberIndex+1])
	isCrowsnest := i.IsCrowsNest(routeNumber)
	isSecondary := i.IsSecondary(routeNumber)

	if isSecondary {
		return SignTemplate{
			Template: "Alberta-Secondary.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     585,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     440,
					Y:            930,
					Center:       469,
				},
			},
		}, true
	} else if isCrowsnest {
		return SignTemplate{
			Template: "Alberta-Crowsnest.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     690,
					DefaultFont:  E,
					OversizeFont: D,
					Text:         tokens[1],
					Color:        i.CrowsnestColor(),
					FontSize:     600,
					Y:            1099,
					Center:       590,
				},
			},
		}, true
	} else {
		return SignTemplate{
			Template: "Alberta.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     740,
					DefaultFont:  C,
					OversizeFont: B,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     630,
					Y:            951,
					Center:       469,
				},
			},
		}, true
	}

}
