package blocklist

func (b blockList) remove(url string) error {
	_, ok := b[url]
	if !ok {
		return doesNotExist
	}
	delete(b, url)
	return nil
}
