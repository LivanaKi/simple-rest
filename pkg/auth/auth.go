package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const (
	AdminUsername = "admin"
	AdminPassword = "password"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		payload, _ := decodeBase64(strings.TrimPrefix(auth, "Basic "))
		parts := strings.SplitN(payload, ":", 2)
		if len(parts) != 2 || parts[0] != AdminUsername || parts[1] != AdminPassword {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func decodeBase64(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	return string(decoded), err
}
