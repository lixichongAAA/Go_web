package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	post := Post{
		Id:      1,
		Content: "Lxc必进字节",
		Author: Author{
			Id:   1,
			Name: "Lxc",
		},
		Comments: []Comment{
			{
				Id:      1,
				Content: "hahaha",
				Author:  "Lxc",
			},
			{
				Id:      2,
				Content: "略略",
				Author:  "Adma",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Marshal Error: ", err)
		return
	}

	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing file ", err)
		return
	}

}
