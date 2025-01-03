package handler

import (
	"net/http"

	"github.com/fatihesergg/go_blog/internal/storage"
	"github.com/fatihesergg/go_blog/internal/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		util.ExecuteTemplate(w, r, "login.html", nil)
		return
	}
	if r.Method == "POST" {

		type ErrorModel struct {
			Msg  []string
			Data interface{}
		}

		var errModel ErrorModel
		err := r.ParseForm()
		if err != nil {
			errModel.Msg = append(errModel.Msg, "Something went wrong")
			util.ExecuteTemplate(w, r, "login.html", errModel)
			return
		}
		email := r.FormValue("email")
		password := r.FormValue("password")

		exist, err := storage.PostgresStore.GetUserByEmail(email)
		if err != nil {
			data := struct {
				Email    string
				Password string
			}{
				Email:    email,
				Password: password,
			}
			if err == gorm.ErrRecordNotFound {
				errModel.Msg = append(errModel.Msg, "Invalid email or password")
			} else {
				errModel.Msg = append(errModel.Msg, "Something went wrong")
			}
			errModel.Data = data
			util.ExecuteTemplate(w, r, "login.html", errModel)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(exist.Password), []byte(password))
		if err != nil {
			data := struct {
				Email    string
				Password string
			}{
				Email:    email,
				Password: password,
			}
			errModel.Msg = append(errModel.Msg, "Invalid email or password")
			errModel.Data = data
			util.ExecuteTemplate(w, r, "login.html", errModel)
			return
		}

		token, err := util.GenerateJsonWebToken(exist.ID)
		if err != nil {
			errModel.Msg = append(errModel.Msg, "Something went wrong")
			util.ExecuteTemplate(w, r, "login.html", errModel)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			HttpOnly: true,
		})
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}
}
