package main

import (
	"login/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// creating router
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/index", controllers.Index).Methods("GET")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/welcome", controllers.Welcome).Methods("GET")
	r.HandleFunc("/logout", controllers.Logout)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
