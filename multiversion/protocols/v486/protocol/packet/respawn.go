package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	RespawnStateSearchingForSpawn = iota
	RespawnStateReadyToSpawn
	RespawnStateClientReadyToSpawn
)

// Respawn is sent by the server to make a player respawn client-side. It is sent in response to a
// PlayerAction packet with ActionType PlayerActionRespawn.
// As of 1.13, the server sends two of these packets with different states, and the client sends one of these
// back in order to complete the respawn.
type Respawn struct {
	Position        mgl32.Vec3
	State           byte
	EntityRuntimeID uint64
}

// ID ...
func (*Respawn) ID() uint32 {
	return IDRespawn
}

// Marshal ...
func (pk *Respawn) Marshal(io protocol.IO) {
	io.Vec3(&pk.Position)
	io.Uint8(&pk.State)
	io.Varuint64(&pk.EntityRuntimeID)
}
