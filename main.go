package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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

var db *sql.DB = setupDB()

func main() {
	splashScreen()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/s/", staticHandler)
	http.ListenAndServe(":8080", nil)
}
