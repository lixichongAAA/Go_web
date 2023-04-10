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

func decode(fileName string) (post Post, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open jsonFile error: ", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Decode error: ", err)
		return
	}
	return
}

func unmarshal(fileName string) (post Post, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Read file error: ", err)
		return
	}

	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
		return
	}
	return
}

func main() {
	post, err := decode("post.json")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("post is: ", post)
}
