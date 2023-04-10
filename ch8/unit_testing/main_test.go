package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Post ID is not the same as post.json", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Post Content is not the same as post.json", post.Content)
	}
}

func TestUnmarshal(t *testing.T) {
	post, err := unmarshal("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Post ID is not the same as post.json", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Post contentt is not the same as post.json", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
