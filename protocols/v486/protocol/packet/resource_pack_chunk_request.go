package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ResourcePackChunkRequest is sent by the client to request a chunk of data from a particular resource pack,
// that it has obtained information about in a ResourcePackDataInfo packet.
type ResourcePackChunkRequest struct {
	UUID       string
	ChunkIndex uint32
}

// ID ...
func (*ResourcePackChunkRequest) ID() uint32 {
	return IDResourcePackChunkRequest
}

// Marshal ...
func (pk *ResourcePackChunkRequest) Marshal(io protocol.IO) {
	io.String(&pk.UUID)
	io.Uint32(&pk.ChunkIndex)
}
