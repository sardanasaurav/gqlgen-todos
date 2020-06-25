package resolvers

import (
	"context"
	"encoding/json"
	"gqlgen-todos/graph/model"
	"gqlgen-todos/graph/resolvers/api-resolver-model"
	"io/ioutil"
	"net/http"
	"time"
)

func (r *queryResolver) Book(ctx context.Context, id *string) (*model.Book, error) {
	idvalue := *id
	req, _ := http.NewRequest("GET", "http://localhost:4001/books/"+idvalue, nil)
	req.Header.Add("Content-Type", "application/json")

	getAuthorClient := &http.Client{Timeout: time.Millisecond * time.Duration(10)}
	res, _ := getAuthorClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var bookData map[string]interface{}
	book := api_resolver_model.BookApi{}
	errResAuthor := json.Unmarshal([]byte(body), &bookData)
	if errResAuthor != nil {

	} else {
		byteBooks, _ := json.Marshal(bookData["data"])
		errUnmarshalBook := json.Unmarshal(byteBooks, &book)
		if errUnmarshalBook != nil {

		}
	}
	authorPointer, _ := r.Author(ctx, &book.Author)

	var gqlBook *model.Book
	gqlBook = &model.Book{}
	gqlBook.Author = []*model.Author{authorPointer}

	gqlBook.ID = book.ID
	gqlBook.Title = book.Title

	r.book = gqlBook

	return r.book, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	req, _ := http.NewRequest("GET", "http://localhost:4001/books", nil)
	req.Header.Add("Content-Type", "application/json")

	getAuthorClient := &http.Client{Timeout: time.Millisecond * time.Duration(10)}
	res, _ := getAuthorClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	var bookData map[string]interface{}
	var books []*api_resolver_model.BookApi

	errResAuthor := json.Unmarshal([]byte(body), &bookData)

	if errResAuthor != nil {

	} else {
		byteBooks, _ := json.Marshal(bookData["data"])
		errUnmarshalBooks := json.Unmarshal(byteBooks, &books)
		if errUnmarshalBooks != nil {
		}
	}
	var gqlBooks []*model.Book
	for _, eachBook := range books {
		var gqlBook *model.Book
		gqlBook = &model.Book{}
		gqlBook.ID = eachBook.ID
		gqlBook.Title = eachBook.Title
		authorPointer, _ := r.Author(ctx, &eachBook.Author)
		gqlBook.Author = []*model.Author{authorPointer}
		gqlBooks = append(gqlBooks, gqlBook)
	}

	r.books = gqlBooks

	return r.books, nil
}
