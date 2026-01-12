package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// var (
// 	files   []string
// 	filesMu sync.Mutex
// )

func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := "uploads"
	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		http.Error(w, "failed to create upload dir", http.StatusInternalServerError)
		return
	}

	path := filepath.Join(uploadDir, header.Filename)

	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	n, err := io.Copy(dst, file)
	if err != nil {
		http.Error(w, "failed to upload", http.StatusInternalServerError)
		return
	}

	// filesMu.Lock()
	// files = append(files, header.Filename)
	// filesMu.Unlock()

	fmt.Fprintf(w, "uploaded %d bytes\n", n)
}
