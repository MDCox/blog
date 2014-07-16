package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[len("/"):] == "" {
		indexHandler(w, r)
	} else {
		postHandler(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer rows.Close()

	var id int
	var title, body, slug string
	var date time.Time
	var posts []Post
	for rows.Next() {
		err := rows.Scan(&id, &title, &body, &date, &slug)
		if err != nil {
			fmt.Printf("%s", err)
		}
		var currPost = Post{ID: id, Title: title, Body: body, Date: date, Slug: slug}
		posts = append(posts, currPost)
	}
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, posts)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/"):]
	p, err := loadPost(slug)
	if err != nil {
		p = &Post{Title: "404", Body: fmt.Sprintf("%s", err)}
	}
	t, _ := template.ParseFiles("views/post.html")
	t.Execute(w, p)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	fileName := fmt.Sprintf("static/%s", r.URL.Path[3:])
	http.ServeFile(w, r, fileName)
}
