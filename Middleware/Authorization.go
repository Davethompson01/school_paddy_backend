package middleware

import (
	"net/http"

	auth "github.com/Davethompson01/School_Paddy_golang/app/Auth"
)

type contextKey string

const ClaimsKey contextKey = "claims"

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims := r.Context().Value(ClaimsKey).(*auth.Claims)

			userRole := claims.Role

			for _, allowed := range allowedRoles {
				if userRole == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}
