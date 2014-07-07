package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"net/http"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=user password=pass dbname=blog_db")
	if err != nil {
		panic(err)
	}
	return db
}

func loadPost(title string) (*Post, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Post{Title: title, Body: body}, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT title FROM posts")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	var title string
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			fmt.Fprintf(w, "error")
		}
		fmt.Fprintf(w, "Title: %s\n", title)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/post/"):]
	p, err := loadPost(title)
	if err != nil {
		p = &Post{Title: title}
	}
	t, _ := template.ParseFiles("views/post.html")
	t.Execute(w, p)
}

var db *sql.DB = setupDB()

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/post/", postHandler)
	http.ListenAndServe(":8080", nil)
}
