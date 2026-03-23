package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const userIDKey ctxKey = "userID"

func FromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDKey).(string)
	return id, ok
}

func JWTMiddleware(next http.Handler) http.Handler {
	secret := []byte(os.Getenv("JWT_SECRET"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const bearer = "Bearer "

		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, bearer) {
			http.Error(w, "authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, bearer)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrTokenUnverifiable
			}

			return secret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		if uid, ok := claims["user_id"].(string); ok {
			ctx := context.WithValue(r.Context(), userIDKey, uid)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		http.Error(w, "user id not found in token", http.StatusUnauthorized)
	})
}
