package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprintln(w, header)
}

func body(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

// func main() {
// 	// server := http.Server{
// 	// 	Addr: "127.0.0.1:8080",
// 	// }
// 	// http.HandleFunc("/headers", headers)
// 	// server.ListenAndServe()

// 	server := http.Server{
// 		Addr: "127.0.0.1:8080",
// 	}
// 	http.HandleFunc("/body", body)
// 	server.ListenAndServe()
// }
