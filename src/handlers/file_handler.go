package handler

import (
	"errors"
	"net/http"

	"biyelap.com/biyelap-core/app/core"
)

// FileHandler handles file specific requests
type FileHandler struct {
}

// PostFile posts new file
func (f *FileHandler) PostFile(w http.ResponseWriter, r *http.Request) {
	// reqMultipartParsed
	rmp := core.ParseMultipartForm(r)
	if rmp == nil {
		core.RespondError(w, errors.New("parse error"), http.StatusInternalServerError)
		return
	}
	url, err := core.FileUpload(rmp)
	if err != nil {
		core.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	core.Respond(w, map[string]string{"url": url})
}
