package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ResourcePackChunkData is sent to the client so that the client can download the resource pack. Each packet
// holds a chunk of the compressed resource pack, of which the size is defined in the ResourcePackDataInfo
// packet sent before.
type ResourcePackChunkData struct {
	UUID       string
	ChunkIndex uint32
	DataOffset uint64
	Data       []byte
}

// ID ...
func (*ResourcePackChunkData) ID() uint32 {
	return IDResourcePackChunkData
}

// Marshal ...
func (pk *ResourcePackChunkData) Marshal(io protocol.IO) {
	io.String(&pk.UUID)
	io.Uint32(&pk.ChunkIndex)
	io.Uint64(&pk.DataOffset)
	io.ByteSlice(&pk.Data)
}
