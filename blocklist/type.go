package blocklist

import "errors"

var (
	alreadyExists = errors.New("error already exists")
	doesNotExist  = errors.New("does not exist")
	inBlockList   = errors.New("item in blocklist")
)

type blockList map[string]bool
type BlocksCli struct {
	BlockItems blockList
}

func New() *BlocksCli {
	return &BlocksCli{
		BlockItems: blockList{},
	}
}
