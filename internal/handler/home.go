package handler

import (
	"net/http"

	"github.com/fatihesergg/go_blog/internal/storage"
	"github.com/fatihesergg/go_blog/internal/util"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := storage.PostgresStore.GetAllPosts()
	if err != nil || len(posts) == 0 {
		util.ExecuteTemplate(w, r, "index.html", nil)
		return
	}

	util.ExecuteTemplate(w, r, "index.html", posts)
}
