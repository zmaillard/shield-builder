package config

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"log"
)

type RoadgeekFont int


const (
	B RoadgeekFont = iota
	C RoadgeekFont = iota
	D RoadgeekFont = iota
	E RoadgeekFont = iota
)

func (f RoadgeekFont) String() string {
	switch f {
	case B:
		return "Roadgeek2005SeriesB.ttf"
	case C:
		return "Roadgeek2005SeriesC.ttf"
	case D:
		return "Roadgeek2005SeriesD.ttf"
	case E:
		return "Roadgeek2005SeriesE.ttf"
	default:
		panic("Invalid enum value")
	}
}

func LoadFont (f RoadgeekFont) (*truetype.Font, error) {

	box := rice.MustFindBox("fonts")

	fontBytes, err := box.Bytes(f.String())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return freetype.ParseFont(fontBytes)

}