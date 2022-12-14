package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go/crud/models"
)

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := []models.Book{}

	if res := h.DB.Find(&books); res.Error != nil {
		fmt.Println(res.Error)
	}
	// write header
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
