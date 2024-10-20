package resolver

import (
	"log"
	"net/http"
)

func (h *Handler) Remove(w http.ResponseWriter, r *http.Request) {
	inputURL, err := validateURL(r)
	if err != nil {
		http.Error(w, "invalid item", http.StatusBadRequest)
		return
	}

	err = h.blocklist.BlockItems.Remove(inputURL)
	if err != nil {
		log.Printf("Failed to remove item: %v", err)
	}
	log.Println("Item removed")
}
