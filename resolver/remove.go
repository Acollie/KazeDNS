package resolver

import (
	"log"
	"net/http"
)

func (h *Handler) Remove(w http.ResponseWriter, r *http.Request) {
	err := h.blocklist.BlockItems.Remove("test")
	if err != nil {
		log.Printf("Failed to remove item: %v", err)
	}
	log.Println("Item removed")
}
