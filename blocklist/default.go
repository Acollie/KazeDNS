package blocklist

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func LoadBlockList(inputURL string) []string {
	content, err := readFileFromURL(inputURL)
	if err != nil {
		return []string{}
	}

	return parseBlockList(content)
}

func parseBlockList(content string) []string {
	lines := strings.Split(content, "\n")
	var blockList []string
	for _, line := range lines {
		if strings.HasPrefix(line, "0.0.0.0") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				blockList = append(blockList, parts[1])
			}
		}
	}
	return blockList
}

func readFileFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
