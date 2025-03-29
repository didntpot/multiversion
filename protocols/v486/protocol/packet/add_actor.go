package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddActor is sent by the server to the client to spawn an entity to the player. It is used for every entity
// except other players, for which the AddPlayer packet is used.
type AddActor struct {
	EntityUniqueID  int64
	EntityRuntimeID uint64
	EntityType      string
	Position        mgl32.Vec3
	Velocity        mgl32.Vec3
	Pitch           float32
	Yaw             float32
	HeadYaw         float32
	Attributes      []protocol.AttributeValue
	EntityMetadata  map[uint32]any
	EntityLinks     []legacyprotocol.EntityLink
}

// ID ...
func (*AddActor) ID() uint32 {
	return IDAddActor
}

// Marshal ...
func (pk *AddActor) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.EntityRuntimeID)
	io.String(&pk.EntityType)
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Float32(&pk.HeadYaw)
	protocol.Slice(io, &pk.Attributes)
	io.EntityMetadata(&pk.EntityMetadata)
	protocol.Slice(io, &pk.EntityLinks)
}
