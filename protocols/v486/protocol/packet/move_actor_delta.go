package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	MoveActorDeltaFlagHasX = 1 << iota
	MoveActorDeltaFlagHasY
	MoveActorDeltaFlagHasZ
	MoveActorDeltaFlagHasRotX
	MoveActorDeltaFlagHasRotY
	MoveActorDeltaFlagHasRotZ
	MoveActorDeltaFlagOnGround
	MoveActorDeltaFlagTeleport
	MoveActorDeltaFlagForceMove
)

// MoveActorDelta is sent by the server to move an entity. The packet is specifically optimised to save as
// much space as possible, by only writing non-zero fields.
// As of 1.16.100, this packet no longer actually contains any deltas.
type MoveActorDelta struct {
	Flags           uint16
	EntityRuntimeID uint64
	Position        mgl32.Vec3
	Rotation        mgl32.Vec3
}

// ID ...
func (*MoveActorDelta) ID() uint32 {
	return IDMoveActorDelta
}

// Marshal ...
func (pk *MoveActorDelta) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint16(&pk.Flags)
	if pk.Flags&MoveActorDeltaFlagHasX != 0 {
		io.Float32(&pk.Position[0])
	}
	if pk.Flags&MoveActorDeltaFlagHasY != 0 {
		io.Float32(&pk.Position[1])
	}
	if pk.Flags&MoveActorDeltaFlagHasZ != 0 {
		io.Float32(&pk.Position[2])
	}
	if pk.Flags&MoveActorDeltaFlagHasRotX != 0 {
		io.ByteFloat(&pk.Rotation[0])
	}
	if pk.Flags&MoveActorDeltaFlagHasRotY != 0 {
		io.ByteFloat(&pk.Rotation[1])
	}
	if pk.Flags&MoveActorDeltaFlagHasRotZ != 0 {
		io.ByteFloat(&pk.Rotation[2])
	}
}
