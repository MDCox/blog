package main

import (
	"net/http"
	"io/ioutil"
	"html/template"
)

type Post struct {
    Title string
    Body  []byte
}

func loadPost(title string) (*Post, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
	    return nil, err
    }
    return &Post{Title: title, Body: body}, nil
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/post/"):]
        p, err := loadPost(title)
	if err != nil {
		p = &Post{Title: title}
	}
	t, _ := template.ParseFiles("app/views/post.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/post/", postHandler)
	http.ListenAndServe(":8080", nil)
}
