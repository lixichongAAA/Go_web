package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("Posts.csv")
	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	Allposts := []Post{
		Post{Id: 1, Content: "LxcYYds", Author: "Lxc"},
		Post{Id: 2, Content: "Lxc IS MY FATHER.", Author: "Lxc's son"},
		Post{Id: 3, Content: "Lxc is my grandFather", Author: "777"},
		Post{Id: 4, Content: "hahaha", Author: "sasa"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range Allposts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("Posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	recorder, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range recorder {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
