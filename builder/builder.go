package builder

import (
	"strings"

	"github.com/minhajuddinkhan75/andela/definitions"
	"github.com/minhajuddinkhan75/andela/models"
)

type PostBuilder interface {
	Posts(models.PostSlice) PostBuilder
	Comments(models.CommentSlice) PostBuilder
	Build() definitions.PostSlice
	BuildCSV() []string
}

func NewPostBuilder() PostBuilder {
	return &postBuilder{}
}

type postBuilder struct {
	posts    models.PostSlice
	comments models.CommentSlice
}

func (b *postBuilder) Posts(posts models.PostSlice) PostBuilder {
	b.posts = posts
	return b
}

func (b *postBuilder) Comments(comments models.CommentSlice) PostBuilder {
	b.comments = comments
	return b
}

//Leaving implementation because of shortage of time
//We can implement this method as well to make use of the same loop
//in order to gain performance
func (b *postBuilder) BuildCSV() []string {

	return []string{}
}

func (b *postBuilder) Build() definitions.PostSlice {

	thePosts := make(definitions.PostSlice, 0)

	for _, post := range b.posts {

		thePost := definitions.Post{
			UserID:   post.UserID,
			PostID:   post.PostID,
			Title:    post.Title,
			Body:     post.Body,
			Comments: "",
		}

		comments := make([]string, 0)
		for _, comment := range b.comments {
			if comment.PostId == post.PostID {
				comments = append(comments, comment.Body)
			}

			thePost.Comments = strings.Join(comments, "|")
		}

		thePosts = append(thePosts, thePost)
	}

	return thePosts
}
