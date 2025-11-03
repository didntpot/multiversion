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
	// BlockVersion is the version of blocks (states) of the game.
	BlockVersion int32 = (1 << 24) | (21 << 16) | (60 << 8)
)
