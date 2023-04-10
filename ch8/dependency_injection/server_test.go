package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/2", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 2 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestPutPost(t *testing.T) {
	mux := http.NewServeMux()
	post := &FakePost{}
	mux.HandleFunc("/post/", handleRequest(post))

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"LxcYYds","author":"Lxc"}`)
	request, _ := http.NewRequest("PUT", "/post/2", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	if post.Content != "LxcYYds" {
		t.Errorf("Content is not correct %v", post.Content)
	}
}
