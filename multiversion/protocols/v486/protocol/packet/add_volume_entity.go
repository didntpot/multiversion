package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddVolumeEntity sends a volume entity's definition and metadata from server to client.
type AddVolumeEntity struct {
	EntityRuntimeID    uint64
	EntityMetadata     map[string]any
	EncodingIdentifier string
	InstanceIdentifier string
	EngineVersion      string
}

// ID ...
func (*AddVolumeEntity) ID() uint32 {
	return IDAddVolumeEntity
}

// Marshal ...
func (pk *AddVolumeEntity) Marshal(io protocol.IO) {
	io.Uint64(&pk.EntityRuntimeID)
	io.NBT(&pk.EntityMetadata, nbt.NetworkLittleEndian)
	io.String(&pk.EncodingIdentifier)
	io.String(&pk.InstanceIdentifier)
	io.String(&pk.EngineVersion)
}
