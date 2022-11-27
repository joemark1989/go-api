package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go/crud/models"
	"github.com/gorilla/mux"
)

func (h *handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// read dynamic id param
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//find book by id
	var book models.Book
	if res := h.DB.First(&book, id); res.Error != nil {
		fmt.Println(res.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
