package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddItemActor is sent by the server to the client to make an item entity show up. It is one of the few
// entities that cannot be sent using the AddActor packet
type AddItemActor struct {
	EntityUniqueID  int64
	EntityRuntimeID uint64
	Item            protocol.ItemInstance
	Position        mgl32.Vec3
	Velocity        mgl32.Vec3
	EntityMetadata  map[uint32]any
	FromFishing     bool
}

// ID ...
func (*AddItemActor) ID() uint32 {
	return IDAddItemActor
}

// Marshal ...
func (pk *AddItemActor) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.EntityRuntimeID)
	io.ItemInstance(&pk.Item)
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.EntityMetadata(&pk.EntityMetadata)
	io.Bool(&pk.FromFishing)
}
