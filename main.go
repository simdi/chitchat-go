package main

import (
	"net/http"

	"github.com/simdi/chitchat-go/controllers"
)

func main() {
	// This is a multiplexer that redirects requests to the right handler.
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", controllers.Index)
	// mux.HandleFunc("/err", err)
	// mux.HandleFunc("/login", login)
	// mux.HandleFunc("/logout", logout)
	// mux.HandleFunc("/sign-up", signUp)
	// mux.HandleFunc("/sign-up_account", signUpAccount)
	// mux.HandleFunc("/authenticate", authenticate)
	// mux.HandleFunc("/thread/new", newThread)
	// mux.HandleFunc("/thread/create", createThread)
	// mux.HandleFunc("/thread/post", postThread)
	// mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    ":8100",
		Handler: mux,
	}
	server.ListenAndServe()
}
