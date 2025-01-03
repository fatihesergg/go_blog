package handler

import (
	"net/http"

	"github.com/fatihesergg/go_blog/internal/storage"
	"github.com/fatihesergg/go_blog/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

func AdminHandler(w http.ResponseWriter, r *http.Request, cliams *jwt.RegisteredClaims) {
	if r.Method == "GET" {
		posts, err := storage.PostgresStore.GetAllPosts()
		if err != nil {
			util.ExecuteTemplate(w, r, "adminDashboard.html", nil)
			return
		}

		util.ExecuteTemplate(w, r, "adminDashboard.html", posts)
	}
}
