package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	delete() (err error)
	update() (er error)
}

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select id, content, author where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statementt := "insert into posts (content, author) values (?, ?)"
	stmt, err := post.Db.Prepare(statementt)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Err()
	return
}

func (post *Post) update() (err error) {
	_, err = post.Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

func (post *Post) delete() (err error) {
	_, err = post.Db.Exec("delete from posts where id = ?", post.Id)
	return
}
