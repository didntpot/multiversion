package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ChunkRadiusUpdated is sent by the server in response to a RequestChunkRadius packet. It defines the chunk
// radius that the server allows the client to have. This may be lower than the chunk radius requested by the
// client in the RequestChunkRadius packet.
type ChunkRadiusUpdated struct {
	ChunkRadius int32
}

// ID ...
func (*ChunkRadiusUpdated) ID() uint32 {
	return IDChunkRadiusUpdated
}

// Marshal ...
func (pk *ChunkRadiusUpdated) Marshal(io protocol.IO) {
	io.Varint32(&pk.ChunkRadius)
}
