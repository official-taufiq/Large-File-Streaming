package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type FileInfo struct {
	Name string `json:"name"`
}

func filesList(w http.ResponseWriter, r *http.Request) {
	// filesMu.Lock()
	// if len(files) == 0 {
	// 	w.Write([]byte("No files uploaded yet\n"))
	// } else {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(files)
	// }
	// filesMu.Unlock()

	entries, err := os.ReadDir("uploads")
	if err != nil {
		http.Error(w, "error reading directory", http.StatusInternalServerError)
		return
	}

	files := make([]FileInfo, 0, len(entries))
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, FileInfo{Name: e.Name()})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)

}
