package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	functions "github.com/pubudu2003060/proxy_project/internal/server/functions"
)

func UserAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userKey := os.Getenv("USER_API_KEY")
		if userKey == "" {
			functions.RespondwithError(w, http.StatusInternalServerError, "api key not found", fmt.Errorf("api key not found"))
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			functions.RespondwithError(w, http.StatusUnauthorized, "missing api key", fmt.Errorf("missing api key"))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "ApiKey" {
			functions.RespondwithError(w, http.StatusUnauthorized, "invalid auth format", fmt.Errorf("invalid auth format"))
			return
		}

		key := parts[1]
		if key != userKey {
			functions.RespondwithError(w, http.StatusForbidden, "forbidden", fmt.Errorf("forbidden"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
