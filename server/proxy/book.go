package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go/crud/models"
)

func (p *Proxy) Books(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book

	resp, err := http.Get("http://127.0.0.1:4000/books")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &books)
	fmt.Println(books)

	if err != nil {
		fmt.Println(err)
	}

	return books, nil
}

// func (p *Proxy) AddBook(ctx context.Context, input models.NewBook) (*models.Book, error) {
// 	var book *models.Book

// 	book.Author = input.Author
// 	book.Desc = input.Desc
// 	book.Title = input.Title

// 	//append to book table
// 	if res := p.DB.Create(&book); res.Error != nil {
// 		fmt.Println(res.Error)
// 	}
// 	return book, nil
// }
