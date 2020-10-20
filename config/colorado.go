package config

import "strings"

type Colorado struct {

}


func (i Colorado) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Colorado) Prefix() string {
	return "CO"
}

func (i Colorado) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}


	// Check Suffix
	suffixTokens := strings.Split(tokens[1], "_")

		return SignTemplate{
			Template: "Colorado.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     340,
					DefaultFont:  E,
					OversizeFont: D,
					Text:         suffixTokens[0],
					Color:        i.Color(),
					FontSize:     183,
					Y:            353,
					Center:       194.8,
				},
			},
		}, true
}
