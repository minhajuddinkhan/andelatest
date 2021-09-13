package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/minhajuddinkhan75/andela/models"
)

var _ PostsRespository = &postRepo{}

type PostsRespository interface {
	GetPosts() (models.PostSlice, error)
}

func NewRepository(postsUrl string) PostsRespository {
	return &postRepo{url: postsUrl}
}

type postRepo struct {
	url string
}

func (r *postRepo) GetPosts() (models.PostSlice, error) {

	request, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts", nil)
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

	var posts []models.Post
	return posts, json.Unmarshal(b, &posts)

}
