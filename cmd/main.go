package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/minhajuddinkhan75/andela/builder"
	repo "github.com/minhajuddinkhan75/andela/repositories"
)

func main() {

	postsRepo := repo.NewRepository("https://jsonplaceholder.typicode.com/posts")
	commentsRepo := repo.NewCommentRepository("https://jsonplaceholder.typicode.com/comments")

	//can use goorutines here for concurrency
	posts, err := postsRepo.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	comments, _ := commentsRepo.GetComments()
	if err != nil {
		log.Fatal(err)
	}

	pb := builder.NewPostBuilder()
	postWithComments := pb.Posts(posts).Comments(comments).Build()

	csvFile, err := os.Create("./data.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, post := range postWithComments {
		var row []string
		row = append(row, fmt.Sprint(post.UserID))
		row = append(row, fmt.Sprint(post.PostID))
		row = append(row, post.Title)
		row = append(row, post.Body)
		row = append(row, post.Comments)
		writer.Write(row)
	}
	writer.Flush()

}
