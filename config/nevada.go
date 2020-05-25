package config

import "strings"

type Nevada struct {

}


func (i Nevada) Color() Color {
	return Color{Blue: 0, Green: 0, Red: 0}
}

func (i Nevada) Prefix() string {
	return "NV"
}

func (i Nevada) Match(pattern string) (SignTemplate, bool) {
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, false
	}

	if strings.ToUpper(tokens[0]) != i.Prefix() {
		return SignTemplate{}, false
	}


		return SignTemplate{
			Template: "Nevada.png",
			TextBlocks: []TextBlock{
				{
					MaxWidth:     380,
					DefaultFont:  D,
					OversizeFont: C,
					Text:         tokens[1],
					Color:        i.Color(),
					FontSize:     250,
					Y:            230,
					Center:       300,
				},
			},
		}, true
}