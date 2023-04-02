package main

import (
	"net/http"
	"text/template"
)

// func process(w http.ResponseWriter, r *http.Request) {
// 	t := template.Must(template.ParseFiles("tmpl.html"))
// 	t.Execute(w, "Hello World!")
// }

// func process(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("tmpl2.html")
// 	rand.Seed(time.Now().Unix())
// 	t.Execute(w, rand.Intn(10) > 5)
// }

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("pipeline.html")
	t.Execute(w, "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
