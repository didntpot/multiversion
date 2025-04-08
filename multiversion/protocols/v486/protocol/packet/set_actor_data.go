package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetActorData is sent by the server to update the entity metadata of an entity. It includes flags such as
// if the entity is on fire, but also properties such as the air it has left until it starts drowning.
type SetActorData struct {
	EntityRuntimeID uint64
	EntityMetadata  map[uint32]any
	Tick            uint64
}

// ID ...
func (*SetActorData) ID() uint32 {
	return IDSetActorData
}

// Marshal ...
func (pk *SetActorData) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.EntityMetadata(&pk.EntityMetadata)
	io.Varuint64(&pk.Tick)
}
