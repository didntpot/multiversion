package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlayerInput is sent by the client to the server when the player is moving but the server does not allow it
// to update its movement using the MovePlayer packet. It includes situations where the player is riding an
// entity like a boat. If this is the case, the packet is sent roughly every tick.
type PlayerInput struct {
	Movement mgl32.Vec2
	Jumping  bool
	Sneaking bool
}

// ID ...
func (*PlayerInput) ID() uint32 {
	return IDPlayerInput
}

// Marshal ...
func (pk *PlayerInput) Marshal(io protocol.IO) {
	io.Vec2(&pk.Movement)
	io.Bool(&pk.Jumping)
	io.Bool(&pk.Sneaking)
}
