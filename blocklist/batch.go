package blocklist

import (
	"errors"
)

type BatchBlock struct {
	urls   []string
	source string
}

func (b blockList) Batch(batch BatchBlock) error {
	for _, item := range batch.urls {
		err := b.add(item)
		switch {
		case errors.Is(err, errors.Is(err, alreadyExists)):
			continue
		default:
			continue
		}
	}
	return nil
}
