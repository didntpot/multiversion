package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	MoveModeNormal = iota
	MoveModeReset
	MoveModeTeleport
	MoveModeRotation
)

const (
	TeleportCauseUnknown = iota
	TeleportCauseProjectile
	TeleportCauseChorusFruit
	TeleportCauseCommand
	TeleportCauseBehaviour
)

// MovePlayer is sent by players to send their movement to the server, and by the server to update the
// movement of player entities to other players.
type MovePlayer struct {
	EntityRuntimeID          uint64
	Position                 mgl32.Vec3
	Pitch                    float32
	Yaw                      float32
	HeadYaw                  float32
	Mode                     byte
	OnGround                 bool
	RiddenEntityRuntimeID    uint64
	TeleportCause            int32
	TeleportSourceEntityType int32
	Tick                     uint64
}

// ID ...
func (*MovePlayer) ID() uint32 {
	return IDMovePlayer
}

// Marshal ...
func (pk *MovePlayer) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Vec3(&pk.Position)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Float32(&pk.HeadYaw)
	io.Uint8(&pk.Mode)
	io.Bool(&pk.OnGround)
	io.Varuint64(&pk.RiddenEntityRuntimeID)
	if pk.Mode == MoveModeTeleport {
		io.Int32(&pk.TeleportCause)
		io.Int32(&pk.TeleportSourceEntityType)
	}
	io.Varuint64(&pk.Tick)
}
