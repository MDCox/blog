package main

import "time"

type Post struct {
	ID    int
	Title string
	Body  string
	Date  time.Time
	Slug  string
}
