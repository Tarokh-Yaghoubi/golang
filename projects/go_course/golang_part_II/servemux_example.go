package main

// ServeMux is the router provided by Go standard lib net/http package
// A  route multiplexer allows you to match different request handlers
// to different URLs served by your server

// Patterns: These are the URL patterns that ServeMux will accept
// The general pattern is of form: [METHOD ]/[HOST]/[PATH]

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello There"))
	})

	router.HandleFunc("POST /hello/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("user_id")
		w.Write([]byte(fmt.Sprintf("hello, %s", userId)))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)

	}
}
