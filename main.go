package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
	"time"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=blog_db sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func loadPost(slug string) (*Post, error) {
	query := fmt.Sprintf("SELECT * FROM posts WHERE slug='%s'", slug)
	var id int
	var title, body string
	var date time.Time
	err := db.QueryRow(query).Scan(&id, &title, &body, &date, &slug)
	if err != nil {
		return nil, err
	}

	return &Post{Title: title, Body: body, Date: date}, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT title FROM posts")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	var title string
	var posts []string
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			fmt.Fprintf(w, "error")
		}
		posts = append(posts, title)
	}
	fmt.Printf("%v", posts)
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w,posts)

}

func postHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/post/"):]
	p, err := loadPost(slug)
	if err != nil {
		p = &Post{Title: "404", Body: fmt.Sprintf("%s",err)}
	}
	fmt.Printf("%v",p)
	t, _ := template.ParseFiles("views/post.html")
	t.Execute(w, p)
}

var db *sql.DB = setupDB()

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/post/", postHandler)
	http.ListenAndServe(":8080", nil)
}
