// ListenAndServe function is used to start a HTTP SERVER
// ListenAndServe listens on the provided address and uses handler to handle requests
// It will use the default serve mux if the provided handler is nil

// http.Handler, this is on interface which has one function, ServeHttp

package main

import (
	"fmt"
	"net/http"
)

// We can define handlers like this, with implementing the ServeHTTP method

// type TarokhHandler struct{}

// func (t *TarokhHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Servus from Tarokh!"))
// }

// type HelloHandler struct{}

// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello There"))
// }

func main() {

	// this is another way to define a handler, without implementing ServeHTTP method separately
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello There"))
	}))
	fmt.Println("running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}

}
