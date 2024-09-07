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
