package definitions

type PostSlice []Post
type Post struct {
	UserID   int    `json:"userId"`
	PostID   int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Comments string `json:"comments"`
}
