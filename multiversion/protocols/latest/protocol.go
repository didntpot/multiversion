package latest

import (
	"github.com/didntpot/multiversion/multiversion/internal/chunk"
	"github.com/didntpot/multiversion/multiversion/mapping"
)

var (
	// blockMapping is the BlockMapping used for translating blocks between versions.
	blockMapping = mapping.NewBlockMapping(blockStateData)
	// NetworkPersistentEncoding is the Encoding used for sending a Chunk over network. It uses NBT, unlike NetworkEncoding.
	NetworkPersistentEncoding = chunk.NewNetworkPersistentEncoding(blockMapping, BlockVersion)
	// BlockPaletteEncoding is the paletteEncoding used for encoding a palette of block states encoded as NBT.
	BlockPaletteEncoding = chunk.NewBlockPaletteEncoding(blockMapping, BlockVersion)
)

// NewItemMapping ...
func NewItemMapping(direct bool) mapping.Item {
	return mapping.NewItemMapping(itemData(), requiredItemList, ItemVersion, direct)
}

// NewBlockMapping ...
func NewBlockMapping() *mapping.DefaultBlockMapping {
	return blockMapping
}
