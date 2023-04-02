package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	daysOfWeek := []string{"Mon", "Tus", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.ExecuteTemplate(w, "tmpl.html", daysOfWeek)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
