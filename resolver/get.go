package resolver

import (
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	items := h.blocklist.BlockItems.Get()
	for _, item := range items {
		w.Write([]byte(item + "\n"))
	}
	return
}
