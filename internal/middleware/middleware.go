package middleware

import (
	"net/http"
	"strconv"

	"github.com/fatihesergg/go_blog/internal/storage"
	"github.com/fatihesergg/go_blog/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

type HandlerFuncWithClaims func(http.ResponseWriter, *http.Request, *jwt.RegisteredClaims)

func CheckLogin(next HandlerFuncWithClaims) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := isAuthenticated(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
			next(w, r, claims)
		}
	}
}

func isAuthenticated(r *http.Request) (*jwt.RegisteredClaims, bool) {
	token, err := r.Cookie("token")
	if err != nil {
		return nil, false
	}

	if token.Value == "" {
		return nil, false
	}

	claims, err := util.ParseJsonWebToken(token.Value)
	if err != nil {
		return nil, false
	}

	userID, err := strconv.Atoi(claims.ID)
	if err != nil {
		return nil, false
	}

	user, err := storage.PostgresStore.GetUserById(userID)
	if err != nil || user.ID == 0 {
		return nil, false
	}

	return claims, true
}
