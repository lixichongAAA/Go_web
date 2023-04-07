package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error open xml file: ", err)
		return
	}
	defer xmlFile.Close()

	xmldata, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error rading xml file: ", err)
		return
	}

	var post Post
	xml.Unmarshal(xmldata, &post)
	fmt.Println(post)
	fmt.Println(post.Xml)
}
