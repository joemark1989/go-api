package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go/crud/models"
	"github.com/gorilla/mux"
)

func (h *handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	//read param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//find book by id
	if res := h.DB.First(&book, id); res.Error != nil {
		fmt.Println(res.Error)
	}
	//delete the book found by ID
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
