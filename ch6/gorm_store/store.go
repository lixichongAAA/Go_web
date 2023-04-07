package main

import (
	"fmt"
	"time"

	//_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int
	CreatedAt time.Time
}

var Db *gorm.DB

// connect to the Db
func init() {
	var err error
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 171,                                                                                   // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post) // {0 Hello World! Sau Sheong [] 0001-01-01 00:00:00 +0000 UTC}

	// Create a post
	Db.Create(&post)
	fmt.Println(post) // {1 Hello World! Sau Sheong [] 2015-04-13 11:38:50.91815604 +0800 SGT}

	// Add a comment
	comment := Comment{Content: "Good post!", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment)

	// Get comments from a post
	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment

	fmt.Println(comments[0]) // {1 Good post! Joe 1 2015-04-13 11:38:50.920377 +0800 SGT}
}
