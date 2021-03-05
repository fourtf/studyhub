package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/fourtf/studyhub/models"

	"github.com/dgrijalva/jwt-go"
)

//JWTVerify verifies the tokens of incoming client request to authed paths
func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenHeader = r.Header.Get("x-access-token") //Grab the token from the header

		tokenHeader = strings.TrimSpace(tokenHeader)

		if tokenHeader == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(tokenHeader, tk.StandardClaims, func(token *jwt.Token) (interface{}, error) {
			//TODO: Random string as token generation secret
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
