package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Read file error: ", err)
		return
	}
	fmt.Println(data)
	fmt.Println(string(data))
	var post Post
	err = json.Unmarshal(data, &post)
	if err != nil {
		fmt.Println("unmarshal error: ", err)
		return
	}
	fmt.Println(post)

}
