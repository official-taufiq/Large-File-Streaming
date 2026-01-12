package main

import "net/http"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		mu.Lock()
		defer mu.Unlock()
		_, ok := apiKeysByKeys[key]
		if !ok {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})
}
