package handlers

import (
	"fmt"
	"github.com/disintegration/imaging"
	"net/http"
	"path"
	"sign-builder/core"

	"github.com/GeertJohan/go.rice"
)

var largeHeight = 50
var smallHeight = 20

func HandleShieldPostQuery(w http.ResponseWriter, r *http.Request) {
	//Make sure fonts are available when script is loaded
	_ = rice.MustFindBox("fonts")

	_ = rice.MustFindBox("templates")

	params := r.URL.Query()
	patterns, ok := params["shield"]
	if !ok || len(patterns) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Count Not Find Parameter Shield")
		return
	}

	pattern := patterns[0]
	img, err := core.Build(pattern)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	smallImage := imaging.Resize(*img, 0, smallHeight, imaging.Linear)
	smallFileName := path.Join("/tmp", pattern+".png")
	smallKey := path.Join("Shields", "20x", pattern+".png")
	err = imaging.Save(smallImage, smallFileName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = core.UploadS3(smallKey, smallFileName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	largeImage := imaging.Resize(*img, 0, largeHeight, imaging.Linear)
	largeFileName := path.Join("/tmp", pattern+".png")
	largeKey := path.Join("Shields", pattern+".png")
	err = imaging.Save(largeImage, largeFileName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}
	err = core.UploadS3(largeKey, largeFileName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(200)
}
