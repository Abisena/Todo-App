package middleware

import (
	"net/http"

	"github.com/gorilla/context"
)

func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := context.GetOk(r, "userRole")
			if !ok {
				http.Error(w, "User role not found", http.StatusUnauthorized)
				return
			}
			
			hasRole := false
			for _, role := range requiredRoles {
				if userRole == role {
					hasRole = true
					break
				}
			}

			if !hasRole {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}