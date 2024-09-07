package resolver

import "net/http"

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	item := h.blocklist.BlockItems.Add("test")
	println(item)
}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	items := h.blocklist.BlockItems.Get()
	println(items)
}
