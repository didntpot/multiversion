package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// UpdateSubChunkBlocks is essentially just UpdateBlock packet, however for a set of blocks in a sub-chunk.
type UpdateSubChunkBlocks struct {
	Position protocol.SubChunkPos
	Blocks   []protocol.BlockChangeEntry
	Extra    []protocol.BlockChangeEntry
}

// ID ...
func (*UpdateSubChunkBlocks) ID() uint32 {
	return IDUpdateSubChunkBlocks
}

// Marshal ...
func (pk *UpdateSubChunkBlocks) Marshal(io protocol.IO) {
	io.SubChunkPos(&pk.Position)
	protocol.Slice(io, &pk.Blocks)
	protocol.Slice(io, &pk.Extra)
}
