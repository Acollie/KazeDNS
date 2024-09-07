package blocklist

func (b blockList) Check(url string) error {
	_, ok := b[url]
	if !ok {
		return inBlockList
	}
	return nil
}

func (b blockList) Get() []string {
	var urls []string
	for url := range b {
		urls = append(urls, url)
	}
	return urls
}
