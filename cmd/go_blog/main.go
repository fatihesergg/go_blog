package main

import (
	"fmt"
	"net/http"

	"github.com/fatihesergg/go_blog/internal/handler"
	"github.com/fatihesergg/go_blog/internal/middleware"
	"github.com/fatihesergg/go_blog/internal/storage"
	"github.com/fatihesergg/go_blog/internal/util"
	"github.com/go-playground/validator/v10"
)

func main() {
	// load templates
	util.LoadTemplates()
	util.Validate = validator.New(validator.WithRequiredStructEnabled())

	// connection string
	dsn := "postgres://fatih:test@localhost:5432/go_blog?sslmode=disable"

	// init db
	storage.SetupDB(dsn)

	// Setup Routes
	mux := http.NewServeMux()

	// Home
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/home", handler.HomeHandler)

	// Login
	mux.HandleFunc("/login", handler.LoginHandler)

	// Post
	mux.HandleFunc("/post/view/{id}", handler.GetPostHandler)
	mux.HandleFunc("/post/search", handler.SearchPostHandler)

	// Admin
	mux.HandleFunc("/admin/dashboard", middleware.CheckLogin(handler.AdminHandler))
	mux.HandleFunc("/admin/post/{id}/delete", middleware.CheckLogin(handler.DeletePostHander))
	mux.HandleFunc("/admin/post/{id}/update", middleware.CheckLogin(handler.UpdatePostHandler))
	mux.HandleFunc("/admin/post/create", middleware.CheckLogin(handler.CreatePostHandler))

	if err := http.ListenAndServe(":3000", mux); err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 3000")
}
