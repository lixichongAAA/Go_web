package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web programming</title></head>
	<body><h1>Hello World</h1></body>
	<html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}

type Post struct {
	name    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		name:    "Lxc",
		Threads: []string{"我旁边抖腿的是个大sb", "纯的一批"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

// func main() {
// 	server := http.Server{
// 		Addr: "127.0.0.1:8080",
// 	}
// 	http.HandleFunc("/write", writeExample)
// 	http.HandleFunc("/writeHeader", writeHeaderExample)
// 	http.HandleFunc("/header", headerExample)
// 	http.HandleFunc("/jsonExample", jsonExample)
// 	server.ListenAndServe()
// }
