package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/avran02/decoplan/files/enum"
	"github.com/avran02/decoplan/files/pb"
)

type AuthMiddleware struct {
	authClient pb.AuthServiceClient
}

func NewAuthMiddleware(client pb.AuthServiceClient) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: client,
	}
}

func (a *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")

		resp, err := a.authClient.ValidateToken(r.Context(), &pb.ValidateTokenRequest{
			AccessToken: accessToken,
		})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), enum.CtxValueUserID, resp.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
