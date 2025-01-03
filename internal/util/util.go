package util

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/csrf"
)

type FormModel struct {
	Msg  []string
	Data interface{}
}

func FormatDate(t time.Time) string {
	return t.Format("2006/01/02 15:04")
}

var Validate *validator.Validate

func ErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "max":
		return fmt.Sprintf("%s must be less than %s", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be greater than %s", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fe.Field())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}

var Templates *template.Template

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	Templates = Templates.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML { return csrf.TemplateField(r) },
		})
	Templates.ExecuteTemplate(w, templateName, data)
}

func LoadTemplates() {
	files := path.Join("internal", "view", "*.html")
	Templates = template.Must(template.ParseGlob(files))
}

func GenerateJsonWebToken(id int) (string, error) {
	userID := strconv.Itoa(id)
	claims := jwt.RegisteredClaims{
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   "user",
		Audience:  jwt.ClaimStrings{"user"},
		Issuer:    "go_blog",
		ID:        userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// NOTE: I'll change this secret key later
	return token.SignedString([]byte("weleavesecretsimplefornow"))
}

func ParseJsonWebToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// NOTE: I'll change this secret key later
		return []byte("weleavesecretsimplefornow"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
