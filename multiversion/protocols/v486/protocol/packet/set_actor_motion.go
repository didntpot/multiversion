package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SetActorMotion is sent by the server to change the client-side velocity of an entity. It is usually used
// in combination with server-side movement calculation.
type SetActorMotion struct {
	EntityRuntimeID uint64
	Velocity        mgl32.Vec3
}

// ID ...
func (*SetActorMotion) ID() uint32 {
	return IDSetActorMotion
}

// Marshal ...
func (pk *SetActorMotion) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Vec3(&pk.Velocity)
}
