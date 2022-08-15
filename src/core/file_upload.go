package core

import (
	"net/http"
	"path/filepath"
)

// FileUpload fetches the file and uploads to s3
func FileUpload(r *http.Request) (string, error) {
	// Retrieve the file from form data
	n := r.Form.Get("name")
	f, h, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	// return header.Filename, nil
	return PutObject(
		BucketName,
		f,
		n+filepath.Ext(h.Filename),
	)
}
