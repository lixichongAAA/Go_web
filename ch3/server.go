package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

// 串联处理器函数
// func hello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello!!"))
// }

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Handler func called - " + name)
// 		h(w, r)
// 	}
// }

// 串联处理器

// type HelloHandler struct{}

// func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello!!")
// }

// func log(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("Handler call - %T111\n", h)
// 		h.ServeHTTP(w, r)
// 	})
// }

// func protect(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("Handler call - %T222\n", h)
// 		h.ServeHTTP(w, r)
// 	})
// }

// 使用第三方 HttpRouter 多路复用器
// func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	fmt.Fprintf(w, "Hello: %s\n", p.ByName("name"))
// }

type Myhandler struct{}

func (m *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// http.HandleFunc("/hello", log(hello))
	// server.ListenAndServe()

	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// hello := HelloHandler{}
	// http.Handle("/hello", protect(log(hello)))
	// server.ListenAndServe()

	// mux := httprouter.New()
	// mux.GET("/hello/:name", hello)

	// server := http.Server{
	// 	Addr:    "127.0.0.1:8080",
	// 	Handler: mux,
	// }
	// server.ListenAndServe()

	handler := Myhandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
