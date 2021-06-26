package config

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"image"
	_ "image/png"
	"strings"
)

var stateSuffix = map[string]string{
	"AB": "Alberta",
	"AL": "Alabama",
	"AK": "Alaska",
	"AZ": "Arizona",
	"AR": "Arkansas",
	"BC": "British Columbia",
	"CA": "California",
	"CO": "Colorado",
	"CT": "Connecticut",
	"DE": "Delaware",
	"FL": "Florida",
	"GA": "Georgia",
	"HI": "Hawaii",
	"ID": "Idaho",
	"IL": "Illinois",
	"IN": "Indiana",
	"IA": "Iowa",
	"KS": "Kansas",
	"KY": "Kentucky",
	"LA": "Louisiana",
	"ME": "Maine",
	"MD": "Maryland",
	"MA": "Massachusetts",
	"MI": "Michigan",
	"MN": "Minnesota",
	"MS": "Mississippi",
	"MO": "Missouri",
	"MT": "Montana",
	"NE": "Nebraska",
	"NV": "Nevada",
	"NH": "New Hampshire",
	"NJ": "New Jersey",
	"NM": "New Mexico",
	"NY": "New York",
	"NC": "North Carolina",
	"ND": "North Dakota",
	"OH": "Ohio",
	"OK": "Oklahoma",
	"OR": "Oregon",
	"PA": "Pennsylvania",
	"RI": "Rhode Island",
	"SC": "South Carolina",
	"SD": "South Dakota",
	"TN": "Tennessee",
	"TX": "Texas",
	"UT": "Utah",
	"VT": "Vermont",
	"VA": "Virginia",
	"WA": "Washington",
	"WV": "West Virginia",
	"WI": "Wisconsin",
	"WY": "Wyoming",
}

var shields = []Shield{
	Alberta{},
	Arizona{},
	BusinessLoop{},
	California{},
	Colorado{},
	Idaho{},
	Interstate{},
	Iowa{},
	Kansas{},
	Minnesota{},
	Montana{},
	MontanaSecondary{},
	Nebraska{},
	Nevada{},
	NewMexico{},
	NorthDakota{},
	Oregon{},
	SouthDakota{},
	US{},
	Utah{},
	Washington{},
	Wyoming{},
}

func LoadSignTemplate(pattern string) (SignTemplate, error) {
	if len(pattern) == 0 {
		return SignTemplate{}, fmt.Errorf("%s Is Invalid Shield Name", pattern)
	}
	tokens := strings.Split(pattern, "-")
	if len(tokens) != 2 {
		return SignTemplate{}, fmt.Errorf("%s Is Invalid Shield Name", pattern)
	}
	for _, v := range shields {
		if signTemplate, ok := v.Match(pattern); ok {
			return signTemplate, nil
		}
	}
	return SignTemplate{}, fmt.Errorf("%s Did Not Match Shield", pattern)
}

type Suffix interface {
	GetSuffix() string
}

type Shield interface {
	Match(pattern string) (SignTemplate, bool)
	Prefix() string
}

type SignTemplate struct {
	Template   string
	TextBlocks []TextBlock
}

type TextBlock struct {
	MaxWidth            int
	DefaultFont         RoadgeekFont
	OversizeFont        RoadgeekFont
	Text                string
	FontSize            float64
	Color               Color
	Y                   int
	Center              float64
	Right               *float64
	ShrinkLastCharacterBy *float64
}

func (t TextBlock) SetColor(c func(int, int, int, int)) {
	c(t.Color.Red, t.Color.Green, t.Color.Blue, 255)
}

func (t TextBlock) Position(width float64) float64 {
	if t.Right != nil {
		return *t.Right - width
	}
	return t.Center - (width / 2)
}

func (t SignTemplate) LoadTemplateImage() (image.Image, string, error) {
	templatePath := rice.MustFindBox("templates")
	reader, _ := templatePath.Open(t.Template)
	defer reader.Close()

	return image.Decode(reader)
}

type Color struct {
	Red   int
	Green int
	Blue  int
}
