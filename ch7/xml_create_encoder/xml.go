package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Lxc必进字节",
		Author: Author{
			Id:   "2",
			Name: "lxc",
		},
	}

	xmlfile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Create XML file error: ", err)
		return
	}
	encoder := xml.NewEncoder(xmlfile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encode: ", err)
		return
	}

}
