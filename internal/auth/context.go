package auth

import (
	"context"
	"net/http"

	"github.com/andersjbe/quin/internal/auth/jwt"
	"github.com/andersjbe/quin/internal/database"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

func Middleware(db *database.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			email, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user, err := db.GetUserByEmail(r.Context(), email)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}


			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *database.User {
	raw, _ := ctx.Value(userCtxKey).(*database.User)
	return raw
}
