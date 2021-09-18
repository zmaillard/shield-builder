package handlers

import (
	"bytes"
	"fmt"
	"image/png"
	"net/http"
	"sign-builder/core"

	"github.com/GeertJohan/go.rice"
)

func HandleShieldQuery(w http.ResponseWriter, r *http.Request) {
	//Make sure fonts are available when script is loaded
	_ = rice.MustFindBox("fonts")

	_ = rice.MustFindBox("templates")

	params := r.URL.Query()
	pattern, ok := params["shield"]
	if !ok || len(pattern) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Count Not Find Parameter Shield")
		return
	}

	img, err := core.Build(pattern[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	var buff bytes.Buffer
	png.Encode(&buff, *img)

	w.WriteHeader(200)
	w.Header().Set("Content-type", "image/png")
	w.Write(buff.Bytes())
}
