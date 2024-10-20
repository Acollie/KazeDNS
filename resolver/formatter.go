package resolver

import (
	"fmt"
	"net/http"
	"net/url"
)

func validateURL(r *http.Request) (string, error) {
	inputURL := r.URL.Query().Get(httpInputQuery)
	if inputURL == "" {
		return "", fmt.Errorf("inputURL is required")
	}

	// parse inputURL
	_, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	return inputURL, nil
}
