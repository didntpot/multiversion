package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddPainting is sent by the server to the client to make a painting entity show up. It is one of the few
// entities that cannot be sent using the AddActor packet.
type AddPainting struct {
	EntityUniqueID  int64
	EntityRuntimeID uint64
	Position        mgl32.Vec3
	Direction       int32
	Title           string
}

// ID ...
func (*AddPainting) ID() uint32 {
	return IDAddPainting
}

// Marshal ...
func (pk *AddPainting) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.EntityRuntimeID)
	io.Vec3(&pk.Position)
	io.Varint32(&pk.Direction)
	io.String(&pk.Title)
}
