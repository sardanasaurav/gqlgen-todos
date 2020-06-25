package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"gqlgen-todos/graph/model"
	"gqlgen-todos/graph/resolvers/api-resolver-model"
	"io/ioutil"

	"net/http"
	"time"
)

func (r *queryResolver) Author(ctx context.Context, id *string) (*model.Author, error) {

	idvalue := *id

	req, _ := http.NewRequest("GET", "http://localhost:4001/authors/"+idvalue, nil)
	req.Header.Add("Content-Type", "application/json")
	//q := req.URL.Query()
	//q.Add("state", "gujarat")
	//req.URL.RawQuery = q.Encode()
	getAuthorClient := &http.Client{Timeout: time.Millisecond * time.Duration(10)}
	res, _ := getAuthorClient.Do(req)
	fmt.Println(res)
	body, _ := ioutil.ReadAll(res.Body)
	var authorData map[string]interface{}
	author := api_resolver_model.AuthorApi{}

	errResAuthor := json.Unmarshal([]byte(body), &authorData)
	if errResAuthor != nil {

	} else {
		byteAuthors, _ := json.Marshal(authorData["data"])
		errUnmarshalAuthor := json.Unmarshal(byteAuthors, &author)
		if errUnmarshalAuthor != nil {

		}

	}

	var gqlAuthor *model.Author

	gqlAuthor = &model.Author{}
	gqlAuthor.ID = author.ID
	gqlAuthor.Name = author.Name
	gqlAuthor.Book = author.Book

	r.author = gqlAuthor

	return r.author, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	req, _ := http.NewRequest("GET", "http://localhost:4001/authors", nil)
	req.Header.Add("Content-Type", "application/json")
	getAuthorClient := &http.Client{Timeout: time.Millisecond * time.Duration(10)}
	res, _ := getAuthorClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var authorData map[string]interface{}
	var author []*model.Author

	errResAuthor := json.Unmarshal([]byte(body), &authorData)
	if errResAuthor != nil {

	} else {
		byteAuthors, _ := json.Marshal(authorData["data"])
		errUnmarshalAuthor := json.Unmarshal(byteAuthors, &author)
		if errUnmarshalAuthor != nil {

		}
	}

	r.authors = author

	return r.authors, nil
}
