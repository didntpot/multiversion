package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// LevelChunk is sent by the server to provide the client with a chunk of a world data (16xYx16 blocks).
// Typically, a certain amount of chunks is sent to the client before sending it the spawn PlayStatus packet,
// so that the client spawns in a loaded world.
type LevelChunk struct {
	Position        protocol.ChunkPos
	HighestSubChunk uint16
	SubChunkCount   uint32
	CacheEnabled    bool
	BlobHashes      []uint64
	RawPayload      []byte
}

// ID ...
func (*LevelChunk) ID() uint32 {
	return IDLevelChunk
}

// Marshal ...
func (pk *LevelChunk) Marshal(io protocol.IO) {
	io.ChunkPos(&pk.Position)
	io.Varuint32(&pk.SubChunkCount)
	if pk.SubChunkCount == protocol.SubChunkRequestModeLimited {
		io.Uint16(&pk.HighestSubChunk)
	}
	io.Bool(&pk.CacheEnabled)
	if pk.CacheEnabled {
		protocol.FuncSlice(io, &pk.BlobHashes, io.Uint64)
	}
	io.ByteSlice(&pk.RawPayload)
}
