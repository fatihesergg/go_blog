package handler

import (
	"net/http"
	"strconv"

	"github.com/fatihesergg/go_blog/internal/model"
	"github.com/fatihesergg/go_blog/internal/storage"
	"github.com/fatihesergg/go_blog/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

func SearchPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		if len(query) < 3 || len(query) > 20 {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		posts, err := storage.PostgresStore.SearchPost(query)
		if err != nil {
			util.ExecuteTemplate(w, r, "searchResult.html", nil)
			return
		}
		formModel := util.FormModel{Msg: []string{query}, Data: posts}
		util.ExecuteTemplate(w, r, "searchResult.html", formModel)
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request, claims *jwt.RegisteredClaims) {
	if r.Method == "GET" {
		util.ExecuteTemplate(w, r, "createPost.html", nil)
		return
	}
	if r.Method == "POST" {

		formModel := util.FormModel{Msg: nil, Data: nil}
		err := r.ParseForm()
		if err != nil {
			formModel.Msg = append(formModel.Msg, "Something went wrong")

			util.ExecuteTemplate(w, r, "createPost.html", formModel)
			return

		}
		title := r.FormValue("title")
		content := r.FormValue("content")

		validate := util.Validate
		if content == "<p><br></p>" {
			content = ""
		}

		err = validate.Struct(model.Post{Title: title, Content: content})
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, v := range validationErrors {
				formModel.Msg = append(formModel.Msg, util.ErrorMsg(v))
			}

			formModel.Data = model.Post{Title: title, Content: content}

			util.ExecuteTemplate(w, r, "createPost.html", formModel)
			return
		}

		post := model.Post{Title: title, Content: content}

		userID, _ := strconv.Atoi(claims.ID)

		user, err := storage.PostgresStore.GetUserById(userID)
		if err != nil {
			formModel.Msg = append(formModel.Msg, "Something went wrong")
			util.ExecuteTemplate(w, r, "createPost.html", formModel)
			return
		}
		post.User = user
		post.UserID = userID

		err = storage.PostgresStore.AddPost(post)
		if err != nil {
			formModel.Msg = append(formModel.Msg, "Something went wrong")
			util.ExecuteTemplate(w, r, "createPost.html", formModel)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		post, err := storage.PostgresStore.GetPostById(id)
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		util.ExecuteTemplate(w, r, "viewPost.html", post)
	}
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request, claims *jwt.RegisteredClaims) {
	formModel := util.FormModel{Msg: nil, Data: nil}
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		post, err := storage.PostgresStore.GetPostById(id)
		if err != nil {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}

		formModel.Data = post
		util.ExecuteTemplate(w, r, "updatePost.html", formModel)
		return
	}
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			formModel.Msg = append(formModel.Msg, "Something went wrong")

			util.ExecuteTemplate(w, r, "updatePost.html", formModel)
			return
		}
		id, _ := strconv.Atoi(r.PathValue("id"))
		title := r.FormValue("title")
		content := r.FormValue("content")
		post := model.Post{Id: id, Title: title, Content: content}

		if title == "" || content == "" {
			formModel.Msg = append(formModel.Msg, "All fields are required")
			formModel.Data = post
			util.ExecuteTemplate(w, r, "updatePost.html", formModel)
			return
		}

		userID, _ := strconv.Atoi(claims.ID)
		post.UserID = userID
		user, _ := storage.PostgresStore.GetUserById(userID)
		post.User = user
		post.UserID = userID
		err = storage.PostgresStore.UpdatePost(post)
		if err != nil {
			formModel.Msg = append(formModel.Msg, "Something went wrong")
			formModel.Data = post

			util.ExecuteTemplate(w, r, "updatePost.html", formModel)
			return
		}
		http.Redirect(w, r, "/post/view/"+strconv.Itoa(id), http.StatusSeeOther)
		return
	}
}

func DeletePostHander(w http.ResponseWriter, r *http.Request, claims *jwt.RegisteredClaims) {
	postID := r.PathValue("id")
	if postID == "" {
		return
	}
	id, err := strconv.Atoi(postID)
	if err != nil {
		return
	}
	_, err = storage.PostgresStore.GetPostById(id)
	if err != nil {
		return
	}
	err = storage.PostgresStore.DeletePost(id)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
