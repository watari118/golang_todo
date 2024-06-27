package controllers

import (
	"log"
	"net/http"

	"local.package/golang_todo/app/models"
)

type PageData struct {
	LoginFailed   bool
	ErrorMessages []string
	User          models.User
	Todo          models.Todo
	Email         string
	Name          string
}

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}

		_, err = models.GetUserByEmail(user.Email)
		if err == nil {
			var errorMessages = []string{"入力したEmailは既に登録されています。"}
			data := PageData{LoginFailed: true, ErrorMessages: errorMessages, Email: user.Email, Name: user.Name}
			generateHTML(w, data, "layout", "public_navbar", "signup")
			return
		}

		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
		return
	}
	http.Redirect(w, r, "/todos", 302)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.PostFormValue("email")
	user, err := models.GetUserByEmail(email)
	if err != nil {
		var errorMessages = []string{"ログインに失敗しました。", "入力したEmailは登録されていません。"}
		data := PageData{LoginFailed: true, ErrorMessages: errorMessages, Email: email}
		generateHTML(w, data, "layout", "public_navbar", "login")
		return
	}
	if user.PassWord != models.Encrypt(r.PostFormValue("password")) {
		var errorMessages = []string{"ログインに失敗しました。", "パスワードを確認してください。"}
		data := PageData{LoginFailed: true, ErrorMessages: errorMessages, Email: email}
		generateHTML(w, data, "layout", "public_navbar", "login")
		return
	}

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.UUID,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}

func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(writer, request, "/login", 302)
}
