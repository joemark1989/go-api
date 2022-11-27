package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go/crud/models"
)

func (h *handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book

	json.Unmarshal(body, &book)

	//append to book table
	if res := h.DB.Create(&book); res.Error != nil {
		fmt.Println(res.Error)
	}
	//send a 201 created res
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
