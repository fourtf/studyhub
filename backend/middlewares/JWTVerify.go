package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/fourtf/studyhub/models"

	"github.com/dgrijalva/jwt-go"
)

//JWTVerify verifies the tokens of incoming client request to authed paths
func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenHeader = r.Header.Get("Token") //Grab the token from the header

		tokenHeader = strings.TrimSpace(tokenHeader)

		if tokenHeader == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			return
		}
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("tokenSigningKey")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), models.UserKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
