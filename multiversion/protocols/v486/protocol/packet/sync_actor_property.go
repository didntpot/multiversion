package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SyncActorProperty is an alternative to synced actor data.
type SyncActorProperty struct {
	PropertyData map[string]any
}

// ID ...
func (*SyncActorProperty) ID() uint32 {
	return IDSyncActorProperty
}

// Marshal ...
func (pk *SyncActorProperty) Marshal(io protocol.IO) {
	io.NBT(&pk.PropertyData, nbt.NetworkLittleEndian)
}
