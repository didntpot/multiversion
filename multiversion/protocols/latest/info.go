package latest

import _ "embed"

var (
	//go:embed assets/required_item_list.json
	requiredItemList []byte
	//go:embed assets/vanilla_items.nbt
	itemRuntimeIDData []byte
	//go:embed assets/block_states.nbt
	blockStateData []byte
)

const (
	// ItemVersion represents the current item version.
	ItemVersion = 191
	// BlockVersion is the version of blocks (states) of the game. This version is composed
	// of 4 (or 3) bytes indicating a version, interpreted as a big endian int. The current version represents
	// 1.21.70 {1, 21, 70}.
	BlockVersion int32 = (1 << 24) | (21 << 16) | (70 << 8)
)
