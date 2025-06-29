package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SubChunk sends data about multiple sub-chunks around a center point.
type SubChunk struct {
	CacheEnabled    bool
	Dimension       int32
	Position        protocol.SubChunkPos
	SubChunkEntries []legacyprotocol.SubChunkEntry
}

// ID ...
func (*SubChunk) ID() uint32 {
	return IDSubChunk
}

// Marshal ...
func (pk *SubChunk) Marshal(io protocol.IO) {
	io.Bool(&pk.CacheEnabled)
	io.Varint32(&pk.Dimension)
	io.SubChunkPos(&pk.Position)
	if pk.CacheEnabled {
		protocol.SliceUint32Length(io, &pk.SubChunkEntries)
	} else {
		protocol.FuncIOSliceUint32Length(io, &pk.SubChunkEntries, legacyprotocol.SubChunkEntryNoCache)
	}
}
