package service

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

type User struct {
	Username string
	Password string
}

func registerHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		if req.Method == "GET" {

			formatter.HTML(w, http.StatusOK, "register", struct{}{})
		} else if req.Method == "POST" {
			err := req.ParseForm()
			if err != nil {

			}

			user := new(User)
			decode := schema.NewDecoder()
			err = decode.Decode(user, req.PostForm)
			if err != nil {

			}

			formatter.HTML(w, http.StatusOK, "inf", struct {
				Username string
				Password string
			}{
				Username: user.Username, Password: user.Password})

		}

	}
}

func jsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{Username: "howlyao", Password: "1234"})
	}
}

func unKnownHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, "501")
	}
}
