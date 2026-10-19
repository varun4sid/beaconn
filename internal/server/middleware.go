package server

import (
	"fmt"
	"net/http"
	"strings"
)

func corsMiddleware(appEnv, allowedDomain string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if originAllowed(appEnv, allowedDomain, origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Add("Vary", "Origin")
			}
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func originAllowed(appEnv, allowedDomain, origin string) bool {
	if origin == "" {
		return false
	}

	if appEnv == "development" &&
		(origin == "http://localhost:4200" || origin == "http://127.0.0.1:4200") {
		return true
	}

	if allowedDomain == origin {
		return true
	}

	if !strings.Contains(allowedDomain, "://") {
		return origin == fmt.Sprintf("http://%s", allowedDomain) ||
			origin == fmt.Sprintf("https://%s", allowedDomain)
	}

	return false
}
