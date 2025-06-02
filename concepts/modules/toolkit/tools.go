package toolkit

import (
	"crypto/rand"
	"errors"
	"net/http"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variable of this type whill have access
// to all the methods with the reciever *Tools.
type Tools struct {
	MaxFileSize int
}

// UploadedFile is a struct used to save information about an uploaded file.
type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int64
}

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string.
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}

func (t *Tools) UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	var uploadedFiles []*UploadedFile

	if t.MaxFileSize <= 0 {
		t.MaxFileSize = 10 * 1024 * 1024 // Default to 10 MB
	}

	err := r.ParseMultipartForm(int64(t.MaxFileSize))
	if err != nil {
		return nil, errors.New("the upload file is too big")
	}

	for _, fHeaders := range r.MultipartForm.File {
		for _, fh := range fHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadFile, error) {
			}(uploadedFiles)
		}
	}

	return uploadedFiles, nil
}
