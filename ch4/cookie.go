package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "Lxc",
		Value:    "YYDS",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "LXC",
		Value:    "kkkk",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func setCookie2(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "Lxc",
		Value:    "YYDS",
		HttpOnly: true,
		MaxAge:   30,
	}
	c2 := http.Cookie{
		Name:     "LXC",
		Value:    "kkkk",
		HttpOnly: true,
		MaxAge:   24,
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c := r.Header["Cookie"]
	fmt.Fprintln(w, c)
}

func getCookie2(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Lxc")
	if err != nil {
		fmt.Fprintf(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c)
	fmt.Fprintln(w, cs)
}

// func main() {
// 	server := http.Server{
// 		Addr: "127.0.0.1:8080",
// 	}
// 	http.HandleFunc("/set_cookie", setCookie)
// 	http.HandleFunc("/set_cookie2", setCookie2)
// 	http.HandleFunc("/get_cookie", getCookie)
// 	http.HandleFunc("/get_cookie2", getCookie2)
// 	server.ListenAndServe()
// }
