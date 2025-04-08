package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	MoveFlagOnGround = 1 << iota
	MoveFlagTeleport
)

// MoveActorAbsolute is sent by the server to move an entity to an absolute position. It is typically used
// for movements where high accuracy isn't needed, such as for long range teleporting.
type MoveActorAbsolute struct {
	EntityRuntimeID uint64
	Flags           byte
	Position        mgl32.Vec3
	Rotation        mgl32.Vec3
}

// ID ...
func (*MoveActorAbsolute) ID() uint32 {
	return IDMoveActorAbsolute
}

// Marshal ...
func (pk *MoveActorAbsolute) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.Flags)
	io.Vec3(&pk.Position)
	io.ByteFloat(&pk.Rotation[0])
	io.ByteFloat(&pk.Rotation[1])
	io.ByteFloat(&pk.Rotation[2])
}
