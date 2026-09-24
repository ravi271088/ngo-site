package middleware

import (
	"net/http"
)

func AdminAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simplified auth for now: check for a basic admin token in header
		// In production, this would be a JWT or Session cookie
		token := r.Header.Get("X-Admin-Token")
		if token != "secret-admin-token-123" {
			http.Error(w, "Unauthorized: Admin access only", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
