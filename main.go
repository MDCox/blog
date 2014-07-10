package main

import (
	"fmt"
//	"log"
	"database/sql"
	_ "github.com/lib/pq"
//	"html/template"
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

func splashScreen() {
	fmt.Printf(" ==============================\n")
	fmt.Printf("|         Golb started         |\n")
	fmt.Printf("|}>-                        -<{|\n")
	fmt.Printf("%v", time.Now())
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

//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	rows, err := db.Query("SELECT * FROM posts")
//	if err != nil {
//		log.Panic(err)
//	}
//	defer rows.Close()
//
//	var id int
//	var title, body, slug string
//	var date time.Time
//	var posts []Post
//	for rows.Next() {
//		err := rows.Scan(&id, &title, &body, &date, &slug)
//		if err != nil {
//			fmt.Fprintf(w, "error")
//		}
//		var currPost = Post{ID: id, Title: title, Date: date, Slug: slug}
//		posts = append(posts, currPost)
//	}
//	t, _ := template.ParseFiles("views/index.html")
//	t.Execute(w,posts)
//
//}

//func postHandler(w http.ResponseWriter, r *http.Request) {
//	slug := r.URL.Path[len("/post/"):]
//	p, err := loadPost(slug)
//	if err != nil {
//		p = &Post{Title: "404", Body: fmt.Sprintf("%s",err)}
//	}
//	t, _ := template.ParseFiles("views/post.html")
//	t.Execute(w, p)
//}

var db *sql.DB = setupDB()

func main() {
	splashScreen()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/p/", postHandler)
	http.ListenAndServe(":8080", nil)
}
