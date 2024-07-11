package blocklist

func (b blockList) Check(url string) error {
	_, ok := b[url]
	if !ok {
		return inBlockList
	}
	return nil
}
