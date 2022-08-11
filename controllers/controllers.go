package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var user string
var store = sessions.NewCookieStore([]byte(user))

func Index(response http.ResponseWriter, request *http.Request) {
	_, err := request.Cookie(user)
	if err != nil {
		tmp, _ := template.ParseFiles("viewPage/index.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/welcome", http.StatusSeeOther)

	}
}

func Login(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	if username == "jashfeer" && password == "123" {
		user = username
		session, _ := store.Get(request, user)
		session.Values["username"] = username
		session.Save(request, response)
		http.Redirect(response, request, "/welcome", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"err": "Invalid username or password",
		}
		tmp, _ := template.ParseFiles("viewPage/index.html")
		tmp.Execute(response, data)
	}

}

func Welcome(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")

	session, _ := store.Get(request, user)
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	if username != nil {
		tmp, _ := template.ParseFiles("viewPage/welcome.html")
		tmp.Execute(response, data)
	} else {
		http.Redirect(response, request, "/index", http.StatusSeeOther)
	}

}
func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, user)
	session.Options.MaxAge = -1
	session.Save(request, response)
	http.Redirect(response, request, "/index", http.StatusSeeOther)
}
