package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// RemoveVolumeEntity indicates a volume entity to be removed from server to client.
type RemoveVolumeEntity struct {
	EntityRuntimeID uint64
}

// ID ...
func (*RemoveVolumeEntity) ID() uint32 {
	return IDRemoveVolumeEntity
}

// Marshal ...
func (pk *RemoveVolumeEntity) Marshal(io protocol.IO) {
	io.Uint64(&pk.EntityRuntimeID)
}
