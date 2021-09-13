package models

type PostSlice []Post

type Post struct {
	UserID int    `json:"userId"`
	PostID int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
