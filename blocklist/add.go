package blocklist

func (b blockList) add(url string) error {
	_, ok := b[url]
	if ok {
		return alreadyExists
	}
	b[url] = true
	return nil
}

func (b blockList) Add(url string) error {
	return b.add(url)
}

func (b blockList) BatchAdd(urls []string) error {
	for _, url := range urls {
		err := b.add(url)
		if err != nil {
			return err
		}
	}
	return nil
}
