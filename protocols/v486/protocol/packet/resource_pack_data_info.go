package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// noinspection SpellCheckingInspection
const (
	ResourcePackTypeAddon = iota + 1
	ResourcePackTypeCached
	ResourcePackTypeCopyProtected
	ResourcePackTypeBehaviour
	ResourcePackTypePersonaPiece
	ResourcePackTypeResources
	ResourcePackTypeSkins
	ResourcePackTypeWorldTemplate
)

// ResourcePackDataInfo is sent by the server to the client to inform the client about the data contained in
// one of the resource packs that are about to be sent.
type ResourcePackDataInfo struct {
	UUID          string
	DataChunkSize uint32
	ChunkCount    uint32
	Size          uint64
	Hash          []byte
	Premium       bool
	PackType      byte
}

// ID ...
func (*ResourcePackDataInfo) ID() uint32 {
	return IDResourcePackDataInfo
}

// Marshal ...
func (pk *ResourcePackDataInfo) Marshal(io protocol.IO) {
	io.String(&pk.UUID)
	io.Uint32(&pk.DataChunkSize)
	io.Uint32(&pk.ChunkCount)
	io.Uint64(&pk.Size)
	io.ByteSlice(&pk.Hash)
	io.Bool(&pk.Premium)
	io.Uint8(&pk.PackType)
}
