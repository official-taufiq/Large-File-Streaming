package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func stream(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("fileName")

	path := filepath.Join("uploads", file)
	src, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error reaching file", http.StatusInternalServerError)
		return
	}
	defer src.Close()

	info, err := src.Stat()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set(
		"Content-Disposition",
		fmt.Sprintf(`attachment; filename="%s"`, file),
	)

	http.ServeContent(w, r, file, info.ModTime(), src)
}
