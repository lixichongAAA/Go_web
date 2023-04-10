package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}

func tearDown() {

}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/2", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 2 {
		t.Error("Cannot retrieve JSON post")
	}

}

func TestPut(t *testing.T) {
	json := strings.NewReader(`"content":"SSP", "autho":"LxcYYDS"`)
	request, _ := http.NewRequest("PUT", "/post/2", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
