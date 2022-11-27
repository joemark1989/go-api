package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go/crud/models"
	"github.com/gorilla/mux"
)

func (h *handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic id param
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)
	var book models.Book

	if res := h.DB.First(&book, id); res.Error != nil {
		fmt.Println(res.Error)
	}

	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc
	book.Title = updatedBook.Title

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
