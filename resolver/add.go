package resolver

import "net/http"

func (h *Handler) AddHttp(w http.ResponseWriter, r *http.Request) {
	inputURL, err := validateURL(r)
	if err != nil {
		http.Error(w, "invalid item", http.StatusBadRequest)
		return
	}
	err = h.blocklist.BlockItems.Add(inputURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Add(item string) error {
	return h.blocklist.BlockItems.Add(item)
}

func (h *Handler) AddBatch(urls []string) error {
	return h.blocklist.BlockItems.BatchAdd(urls)
}
