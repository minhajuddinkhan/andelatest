package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/minhajuddinkhan75/andela/models"
)

var _ CommentsRespository = &commentRepo{}

type CommentsRespository interface {
	GetComments() (models.CommentSlice, error)
}

func NewCommentRepository(commentsUrl string) CommentsRespository {
	return &commentRepo{url: commentsUrl}
}

type commentRepo struct {
	url string
}

func (r *commentRepo) GetComments() (models.CommentSlice, error) {

	request, err := http.NewRequest(http.MethodGet, r.url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err

	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("something went wrong")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	var comments []models.Comment
	return comments, json.Unmarshal(b, &comments)

}
