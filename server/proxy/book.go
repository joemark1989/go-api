package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go/crud/models"
)

func (p *Proxy) Books(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book
	fmt.Println(books)

	resp, err := http.Get("http://127.0.0.1:4000/books")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &books)

	if err != nil {
		fmt.Println(err)
	}

	return books, nil
}

func (p *Proxy) AddBook(ctx context.Context, input models.NewBook) (*models.Book, error) {
	book := &models.Book{
		Title:  input.Title,
		Author: input.Author,
		Desc:   input.Desc,
	}
	fmt.Println(book)
	data, _ := json.Marshal(book)

	req, err := http.NewRequest("POST", "http://127.0.0.1:4000/books", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
	}

	//append to book table
	c := &http.Client{}

	res, err := c.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))

	return book, nil
}
