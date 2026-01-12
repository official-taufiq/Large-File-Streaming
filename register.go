package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type reqParams struct {
	Email string `json:"email"`
}

type resParams struct {
	Email  string `json:"email"`
	APIkey string `json:"apikey"`
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
func generateAPIKey() string {
	return fmt.Sprintf("sk-%s", hex.EncodeToString(randBytes(32)))

}

var (
	apiKeysByKeys  = make(map[string]string)
	apiKeysByEmail = make(map[string]string)
	mu             sync.Mutex
)

func register(w http.ResponseWriter, r *http.Request) {
	var req reqParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, ok := apiKeysByEmail[req.Email]
	if !ok {
		key := generateAPIKey()
		apiKeysByEmail[req.Email] = key
		apiKeysByKeys[key] = req.Email
		res := &resParams{
			Email:  req.Email,
			APIkey: key,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	http.Error(w, "user already registered", http.StatusConflict)
}
