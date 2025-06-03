package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ThomasCode92/toolkit"
)

func main() {
	mux := routes()

	log.Println("Starting server on port 8080...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/upload", uploadFiles)
	mux.HandleFunc("/upload-one", uploadFile)

	return mux
}

func uploadFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	t := toolkit.Tools{
		MaxFileSize:      10 * 1024 * 1024, // 10 MB
		AllowedFileTypes: []string{"image/jpeg", "image/png", "image/gif"},
	}

	files, err := t.UploadFiles(r, "./uploads", true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	out := ""
	for _, file := range files {
		out += fmt.Sprintf("Uploaded %s to uploads folder, renamed to %s\n", file.OriginalFileName, file.NewFileName)
	}

	_, _ = w.Write([]byte(out))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
}
